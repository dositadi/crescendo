package ticketpage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	ticketpage "github.com/dositadi/groupie-tracker/internal/services/pages/ticket_page"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/google/uuid"
)

const (
	sourceSold = "Sold Ticket Handler under ticketpage pkg"
)

func (t *TicketPage) SoldTicketHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<10) // 500kb

	if err := r.ParseForm(); err != nil {
		e := helper.WrapError("Byte too large", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceSold,
		})
		http.Error(w, e.Error(), http.StatusRequestEntityTooLarge)
	}

	user := t.getUserId(r)

	var soldTicket data.SoldTickets

	// Cache access data
	location := r.FormValue(utils.LOCATION_KEY)
	date := r.FormValue(utils.DATE_KEY)
	artistId := t.atoi(r.FormValue(utils.ARTIST_ID_KEY), sourceSold)

	// User contact info
	userFirstName := r.FormValue(utils.USER_FN_KEY)
	userLastName := r.FormValue(utils.USER_LN_KEY)
	userEmail := r.FormValue(utils.EMAIL_KEY)

	// Card details
	cardNumber := r.FormValue(utils.CARD_NO_KEY)
	cardExpDate := r.FormValue(utils.CARD_EXP_DATE_KEY)
	cvc := r.FormValue(utils.CVC_KEY)
	cardHolderName := r.FormValue(utils.CARD_HOLDER_NAME_KEY)

	fmt.Printf("Card number: %s\nCard expiry date: %s\nCVC: %s\nCard holder name: %s\n", cardNumber, cardExpDate, cvc, cardHolderName)

	order, exists := ordercache.Get(user.Id, location, artistId)
	if !exists {
		e := errors.New("Order does not exist")
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceSold,
		})

		detailPagePath := fmt.Sprintf("%s/%v", utils.ARTIST_DETAILS.String(), artistId)

		url := fmt.Sprintf("%s?%s=%s&%s=%v&%s=%s&%s=%s", utils.TICKET.String(), utils.DATE_KEY, date, utils.ARTIST_ID_KEY, artistId, utils.LOCATION_KEY, location, utils.PATH_KEY, detailPagePath)

		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	vat := float64(int(ordercache.VAT))
	bookingFee := float64(ordercache.BOOKING_FEE)

	soldTicket = data.SoldTickets{
		Id:               uuid.NewString(),
		UserId:           user.Id,
		UserContactFName: userFirstName,
		UserContactLName: userLastName,
		UserContactEmail: userEmail,
		ArtistId:         artistId,
		ConcertDate:      date,
		TicketType:       order.TicketType,
		Qty:              order.Quantity,
		Vat:              &vat,
		Amt:              &order.GrandTotalAmount,
		Location:         location,
		BookingFee:       &bookingFee,
	}

	if err := t.soldTicketsModel.Insert(soldTicket); err != nil {
		e := helper.WrapError("Sold ticket insertion error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceSold,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	page := ticketpage.New(t.logger, w, t.embedded, t.client, r, t.soldTicketsModel)

	if err := page.RenderPaymentPage(false); err != nil {
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceSold,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
