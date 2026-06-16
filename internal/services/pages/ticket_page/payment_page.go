package ticketpage

import (
	"html/template"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourcePP = "Render payment page under ticketpage pkg"
)

func (t *TicketPage) RenderPaymentPage() error {
	fs := []string{
		"internal/web/static/partials/pages/payment_page.html",
	}

	artistId := t.atoi(t.request.FormValue(utils.ARTIST_ID_KEY))
	location := t.request.FormValue(utils.LOCATION_KEY)
	date := t.request.FormValue(utils.DATE_KEY)
	user := t.getUser()

	booking, exists := ordercache.Get(user.Id, location, artistId)
	if !exists {
		t.logger.PrintError(NOT_FOUND.Error(), map[string]string{
			"Source": sourcePP,
		})
		return NOT_FOUND
	}

	artistInfo := t.client.GetByIdKey()[artistId]

	data := struct {
		ArtistInfo              artistapi.ArtistInfo
		ArtistId, VatValue      int
		Location, Date          string
		Quantity                int
		TicketType              string
		BookingFee, TicketPrice float64
	}{
		Date:        date,
		BookingFee:  float64(ordercache.BOOKING_FEE),
		TicketPrice: ordercache.GetTicketPrice(booking.TicketType),
		VatValue:    int(ordercache.VAT),
		ArtistInfo:  artistInfo,
		ArtistId:    artistId,
		Location:    location,
		Quantity:    booking.Quantity,
		TicketType:  booking.TicketType,
	}

	temp, err := template.New("payment_page.html").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template creation error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourcePa,
		})
		return e
	}

	if err := temp.ExecuteTemplate(t.responseWriter, "payment", data); err != nil {
		e := helper.WrapError("Template execution error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourcePa,
		})
		return e
	}
	return nil
}
