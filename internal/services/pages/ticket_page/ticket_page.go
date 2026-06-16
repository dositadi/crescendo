package ticketpage

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceT = "Render ticket page f(n) under ticketpage pkg"
)

func (t *TicketPage) RenderTicketPage() error {
	fs := []string{
		"internal/web/static/pages/ticket_purchase_page.html",
	}

	artistId := t.atoi(t.request.FormValue(utils.ARTIST_ID_KEY))
	date := t.request.FormValue(utils.DATE_KEY)
	location := t.request.FormValue(utils.LOCATION_KEY)
	user := t.getUser()
	path := fmt.Sprintf("%s/%v", utils.ARTIST_DETAILS.String(), artistId)

	// Add the user's order to the cache
	ordercache.Set(user.Id, location, artistId, string(ordercache.GENERAL))

	var artistInfo artistapi.ArtistInfo

	if val, ok := t.client.GetByIdKey()[artistId]; ok {
		artistInfo = val
	} else {
		err := errors.New("Artist ID does not exist")
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceT,
		})
		http.Error(t.responseWriter, err.Error(), http.StatusBadRequest)
	}

	//Get User's order from cache
	booking, exist := ordercache.Get(user.Id, location, artistId)
	if !exist {
		t.logger.PrintError(NOT_FOUND.Error(), map[string]string{
			"Source": sourcePa,
		})

		detailPagePath := fmt.Sprintf("%s/%v", utils.ARTIST_DETAILS.String(), artistId)

		url := fmt.Sprintf("%s?%s=%s&%s=%v&%s=%s&%s=%s", utils.TICKET.String(), utils.DATE_KEY, date, utils.ARTIST_ID_KEY, artistId, utils.LOCATION_KEY, location, utils.PATH_KEY, detailPagePath)

		http.Redirect(t.responseWriter, t.request, url, http.StatusSeeOther)
	}

	data := struct {
		TicketType                                                                         string
		Quantity                                                                           int
		TicketPrice                                                                        float64
		BookingFee, TotalBookingFee, TotalTicketAmount, TotalVatAmount, GrandTotal         float64
		VatValue                                                                           float64
		ArtistInfo                                                                         artistapi.ArtistInfo
		Location, Date                                                                     string
		PreviousPageUrl, TicketTypeUrl, TicketQtyUrl, PaymentUrl                           string
		ArtistId                                                                           int
		ArtistIdKey, DateKey, LocationKey, TicketTypeKey, IncrementQtyKey, DecrementQtyKey string
		GeneralTicket, VipTicket, ReserveTicket                                            string
		GeneralTicketPrice, VipTicketPrice, ReserveTicketPrice                             float64
	}{
		// User order details from cache
		TicketType:        booking.TicketType,
		Quantity:          booking.Quantity,
		TotalBookingFee:   booking.TotalBookingFee,
		TotalTicketAmount: booking.TotalTicketAmount,
		TotalVatAmount:    booking.TotalVatAmount,
		GrandTotal:        booking.GrandTotalAmount,

		// Tickets
		GeneralTicket: string(ordercache.GENERAL),
		VipTicket:     string(ordercache.VIP),
		ReserveTicket: string(ordercache.RESERVED),

		//Ticket prices
		GeneralTicketPrice: float64(ordercache.GENERAL_AMT),
		VipTicketPrice:     float64(ordercache.VIP_AMT),
		ReserveTicketPrice: float64(ordercache.RESERVED_AMT),
		TicketPrice:        ordercache.GetTicketPrice(booking.TicketType),
		BookingFee:         float64(ordercache.BOOKING_FEE),
		VatValue:           ordercache.Round(float64(ordercache.VAT)),

		// All keys
		TicketTypeKey:   utils.TICKET_TYPE_KEY,
		ArtistIdKey:     utils.ARTIST_ID_KEY,
		DateKey:         utils.DATE_KEY,
		LocationKey:     utils.LOCATION_KEY,
		DecrementQtyKey: utils.DECREMENT_QTY_KEY,
		IncrementQtyKey: utils.INCREMENT_QTY_KEY,

		// Page Urls
		TicketTypeUrl:   utils.TicketType.String(),
		TicketQtyUrl:    utils.TicketQuantity.String(),
		PaymentUrl:      utils.Payment.String(),
		PreviousPageUrl: path,

		// Artist info
		ArtistId:   artistId,
		ArtistInfo: artistInfo,
		Location:   location,
		Date:       date,
	}

	temp, err := template.New("ticket_purchase_page.html").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template parse error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceT,
		})
		return e
	}

	if err := temp.Execute(t.responseWriter, data); err != nil {
		e := helper.WrapError("Template execution error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceT,
		})
		return e
	}
	return nil
}
