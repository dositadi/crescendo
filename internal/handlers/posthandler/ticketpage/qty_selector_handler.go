package ticketpage

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	ticketpage "github.com/dositadi/groupie-tracker/internal/services/pages/ticket_page"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceQ = "Quantity selector handler f(n) under ticketpage pkg"
)

func (t *TicketPage) QuantitySelectorHandler(w http.ResponseWriter, r *http.Request) {
	artistId := t.atoi(r.FormValue(utils.ARTIST_ID_KEY), sourceQ)
	location := r.FormValue(utils.LOCATION_KEY)
	quantityIncrement := r.FormValue(utils.INCREMENT_QTY_KEY)
	quantityDecrement := r.FormValue(utils.DECREMENT_QTY_KEY)
	user := t.getUserId(r)

	if quantityDecrement != "" && quantityDecrement == utils.DECREMENT_QTY_KEY {
		done := ordercache.Decrement(user.Id, location, artistId)
		if !done {
			ordercache.Set(user.Id, location, artistId, string(ordercache.GENERAL))
		}
	} else if quantityIncrement != "" && quantityIncrement == utils.INCREMENT_QTY_KEY {
		done := ordercache.Increment(user.Id, location, artistId)
		if !done {
			ordercache.Set(user.Id, location, artistId, string(ordercache.GENERAL))
		}
	}

	partial := ticketpage.New(t.logger, w, t.embedded, t.client, r,t.soldTicketsModel)

	if err := partial.RenderTicketPagePartials(user.Id, location, artistId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceQ,
		})
	}
}
