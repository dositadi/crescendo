package ticketpage

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceTI = "ticketpage.RenderTicketPage()"
)

func (t *TicketPage) RenderTicketReciept() error {
	fs := []string{
		"internal/web/static/pages/reciepts_page.html",
	}
	artistId := t.request.FormValue(utils.ARTIST_ID_KEY)
	location := t.request.FormValue(utils.LOCATION_KEY)
	date := t.request.FormValue(utils.DATE_KEY)

	user := t.getUser()

	userTicket, err := t.soldTicketsModel.Get(t.atoi(artistId), user.Id, location, date)
	if err != nil {
		e := helper.WrapError("Get ticket error", err)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceTI,
		})
		return e
	}

	data := struct {
		BoughtTicket data.SoldTickets
	}{
		BoughtTicket: userTicket,
	}

	temp, err := template.New("reciepts_page.html").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template parse error", err)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceTI,
		})
		return e
	}

	if err := temp.Execute(t.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceTI,
		})
		return e
	}
	return nil
}
