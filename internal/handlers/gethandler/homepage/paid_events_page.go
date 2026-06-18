package homepage

import (
	"net/http"

	ticketpage "github.com/dositadi/groupie-tracker/internal/services/pages/ticket_page"
)

const (
	sourcePE = "homepage.PaidEventsPageHandler()"
)

func (h *HomePage) PaidEventsPageHandler(w http.ResponseWriter, r *http.Request) {
	pages := ticketpage.New(h.logger, w, h.embedded, h.client, r, h.soldTickets)

	if err := pages.RenderPaidEventsPage(ticketpage.AllEvents); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.PrintError(err.Error(), map[string]string{
			"Source": sourcePE,
		})
	}
}
