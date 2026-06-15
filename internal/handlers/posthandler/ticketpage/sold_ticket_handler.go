package ticketpage

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
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

	location := r.FormValue(utils.LOCATION_KEY)
	date := r.FormValue(utils.DATE_KEY)
	artistId := t.atoi(r.FormValue(utils.ARTIST_ID_KEY), sourceSold)
	total := r.FormValue(utils.TOTAL_AMT_KEY)
	userFirstName := r.FormValue("")
	userLastName := r.FormValue("")
	userEmail := r.FormValue("")

	order, exists := ordercache.Get(user.Id, location, artistId)
	if !exists {
		e := errors.New("Order does not exist")
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceSold,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	vat := float64(int(ordercache.VAT))
	amt, err := strconv.ParseFloat(total, 64)
	if err != nil {
		e := errors.New("Bad total value format (expected float)")
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceSold,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}
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
		Amt:              &amt,
		Location:         location,
		BookingFee:       &bookingFee,
	}

	if err = t.soldTicketsModel.Insert(soldTicket); err != nil {
		e := helper.WrapError("Sold ticket insertion error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceSold,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

}
