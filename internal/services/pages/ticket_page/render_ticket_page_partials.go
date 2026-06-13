package ticketpage

import (
	"errors"
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourcePa = "Render Ticket page partial f(n) under ticketpage pkg"
)

var NOT_FOUND = errors.New("Order does not exist in cache")

func (t *TicketPage) RenderTicketPagePartials(userId, location string, artistId int) error {
	fs := []string{
		"internal/web/static/partials/pages/ticket_page_partials.html",
	}

	booking, exist := ordercache.Get(userId, location, artistId)
	if !exist {
		t.logger.PrintError(NOT_FOUND.Error(), map[string]string{
			"Source": sourcePa,
		})
		return NOT_FOUND
	}

	data := struct {
		TicketType                                                 string
		Quantity                                                   int
		TicketPrice                                                float64
		BookingFee                                                 float64
		VatValue                                                   int
		TicketQtyUrl                                               string
		IncrementQtyKey, DecrementQtyKey, LocationKey, ArtistIdKey string
		ArtistId                                                   int
		Location                                                   string
	}{
		IncrementQtyKey: utils.INCREMENT_QTY_KEY,
		DecrementQtyKey: utils.DECREMENT_QTY_KEY,
		LocationKey:     utils.LOCATION_KEY,
		ArtistIdKey:     utils.ARTIST_ID_KEY,
		ArtistId:        artistId,
		Location:        location,
		TicketQtyUrl:    utils.TicketQuantity.String(),
		TicketType:      booking.TicketType,
		Quantity:        booking.Quantity,
		TicketPrice:     t.getTicketPrice(booking.TicketType),
		BookingFee:      float64(ordercache.BOOKING_FEE),
		VatValue:        int(ordercache.VAT),
	}

	temp := template.Must(template.New("ticket_page_partials.htmls").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...))

	if err := temp.ExecuteTemplate(t.responseWriter, "order-summary", data); err != nil {
		e := helper.WrapError("Template execution error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourcePa,
		})
		return e
	}
	return nil
}
