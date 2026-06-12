package ticketpage

import (
	"errors"
	"html/template"
	"net/http"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceT = "Render ticket page f(n) under ticketpage pkg"
)

func (t *TicketPage) RenderTicketPage() error {
	fs := []string{
		"internal/web/static/pages/ticket_purchase_page.html",
	}

	path := t.request.FormValue(utils.PATH_KEY)
	artistId := t.atoi(t.request.FormValue(utils.ARTIST_ID_KEY))
	date := t.request.FormValue(utils.DATE_KEY)
	location := t.request.FormValue(utils.LOCATION_KEY)

	var artistInfo artistapi.ArtistInfo

	if val, ok := t.client.GetByIdKey()[artistId]; ok {
		artistInfo = val
	} else {
		err := errors.New("Artist ID does not exist")
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceT,
		})
		http.Error(t.responseWriter, err.Error(), http.StatusBadRequest)
	}

	data := struct {
		ArtistInfo      artistapi.ArtistInfo
		Location, Date  string
		PreviousPageUrl string
	}{
		ArtistInfo:      artistInfo,
		Location:        location,
		Date:            date,
		PreviousPageUrl: path,
	}

	temp := template.Must(template.New("ticket_purchase_page.html").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...))

	if err := temp.Execute(t.responseWriter, data); err != nil {
		e := helper.WrapError("Template execution error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceT,
		})
		return e
	}
	return nil
}
