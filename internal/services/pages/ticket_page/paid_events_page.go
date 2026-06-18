package ticketpage

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
)

type Events string

const (
	AllEvents    Events = Events("All events")
	EventReciept Events = Events("Event Reciept")
	sourcePE            = "ticketpage.RenderPaidEventsPage()"
)

func (t *TicketPage) RenderPaidEventsPage(event Events) error {
	fs := []string{
		"internal/web/static/pages/paid_events.html",
	}

	user := t.getUser()

	boughtTickets, err := t.soldTicketsModel.GetAll(user.Id)
	if err != nil {
		t.logger.PrintError(NOT_FOUND.Error(), map[string]string{
			"Source": sourcePE,
		})
		return err
	}

	data := struct {
		PaidEvents []data.SoldTickets
	}{
		PaidEvents: boughtTickets,
	}

	temp, err := template.New("paid_events.html").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template parse error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourcePE,
		})
		return e
	}

	switch event {
	case AllEvents:
		err = temp.Execute(t.responseWriter, data)
	case EventReciept:

	}

	if err != nil {
		e := helper.WrapError("Template execution error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourcePE,
		})
		return e
	}

	return nil
}
