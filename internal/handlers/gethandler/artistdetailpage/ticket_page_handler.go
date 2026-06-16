package artistdetailpage

import (
	"net/http"

	ticketpage "github.com/dositadi/groupie-tracker/internal/services/pages/ticket_page"
)

const (
	sourceT = "Ticket page handler f(n) under artistdetailpage pkg"
)

func (d *DetailPage) TicketPageHandler(w http.ResponseWriter, r *http.Request) {
	pages := ticketpage.New(d.logger, w, d.embedded, d.client, r, nil)

	if err := pages.RenderTicketPage(); err != nil {
		d.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceD,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
