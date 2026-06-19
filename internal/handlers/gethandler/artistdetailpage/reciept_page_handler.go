package artistdetailpage

import (
	"net/http"

	ticketpage "github.com/dositadi/groupie-tracker/internal/services/pages/ticket_page"
)

const (
	sourceR = "artistdetailpage.RecieptPageHandler()"
)

func (a *DetailPage) RecieptPageHandler(w http.ResponseWriter, r *http.Request) {
	page := ticketpage.New(a.logger, w, a.embedded, a.client, r, a.soldTickets)

	if err := page.RenderTicketReciept(); err != nil {
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceR,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
