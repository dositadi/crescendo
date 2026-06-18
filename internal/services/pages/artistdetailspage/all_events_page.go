package artistdetail

import (
	"fmt"
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/go-chi/chi/v5"
)

const (
	sourceE = "Render all events page f(n) under artistdetail pkg"
)

func (a *ArtistDetail) RenderAllEventsPage() error {
	fs := []string{
		"internal/web/static/pages/all_concerts.html",
	}

	id := a.atoi(chi.URLParam(a.request, "id"))

	artistInfo := a.client.Get()[id]

	temp, err := template.New("all_concerts.html").Funcs(a.detailPageFuncMap()).ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template create error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Status": sourceE,
		})
	}
	prevPage := fmt.Sprintf("%s/%v", utils.ARTIST_DETAILS.String(), id)

	data := struct {
		ArtistDetailUrl, PreviousPageUrl, TicketUrl, PathUrl string
		ArtistInfo                                           herokuapp.ArtistInfo
		AllArtists                                           map[int]herokuapp.ArtistInfo
		PathKey, DateKey, ArtistIdKey, LocationKey           string
	}{
		// All urls
		PreviousPageUrl: prevPage,
		ArtistDetailUrl: utils.ARTIST_DETAILS.String(),
		TicketUrl:       utils.TICKET.String(),

		// Artists details
		ArtistInfo: artistInfo,
		AllArtists: a.client.Get(),

		// All keys
		PathKey:     utils.PATH_KEY,
		DateKey:     utils.DATE_KEY,
		ArtistIdKey: utils.ARTIST_ID_KEY,
		LocationKey: utils.LOCATION_KEY,
	}

	if err = temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Status": sourceE,
		})
	}
	return nil
}
