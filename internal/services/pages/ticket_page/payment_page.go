package ticketpage

import (
	"fmt"
	"html/template"
	"net/http"

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

		detailPagePath := fmt.Sprintf("%s/%v", utils.ARTIST_DETAILS.String(), artistId)

		url := fmt.Sprintf("%s?%s=%s&%s=%v&%s=%s&%s=%s", utils.TICKET.String(), utils.DATE_KEY, date, utils.ARTIST_ID_KEY, artistId, utils.LOCATION_KEY, location, utils.PATH_KEY, detailPagePath)

		http.Redirect(t.responseWriter, t.request, url, http.StatusSeeOther)
	}

	artistInfo := t.client.GetByIdKey()[artistId]

	data := struct {
		ArtistInfo                                                                                                     artistapi.ArtistInfo
		ArtistId, VatValue                                                                                             int
		Location, Date                                                                                                 string
		Booking                                                                                                        ordercache.Booking
		BookingFee, TicketPrice                                                                                        float64
		FnKey, LnKey, EmailKey, CardNoKey, ExpiryDateKey, CvcKey, CardHolderNameKey, ArtistIdKey, LocationKey, DateKey string
		PaymentUrl                                                                                                     string
	}{
		// Booking details
		Date:        date,
		Booking:     booking,
		BookingFee:  float64(ordercache.BOOKING_FEE),
		TicketPrice: ordercache.GetTicketPrice(booking.TicketType),
		VatValue:    int(ordercache.VAT),
		Location:    location,

		// Artist details
		ArtistInfo: artistInfo,
		ArtistId:   artistId,

		// All keys
		FnKey:             utils.USER_FN_KEY,
		LnKey:             utils.USER_LN_KEY,
		EmailKey:          utils.EMAIL_KEY,
		CardNoKey:         utils.CARD_NO_KEY,
		ExpiryDateKey:     utils.CARD_EXP_DATE_KEY,
		CvcKey:            utils.CVC_KEY,
		CardHolderNameKey: utils.CARD_HOLDER_NAME_KEY,
		ArtistIdKey:       utils.ARTIST_ID_KEY,
		LocationKey:       utils.LOCATION_KEY,
		DateKey:           utils.DATE_KEY,

		//Url
		PaymentUrl: utils.Payment.String(),
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
