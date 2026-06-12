package artistdetail

import (
	"html/template"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/go-chi/chi/v5"
)

const (
	sourceR = "Render artist detail page f(n) under artistdetail pkg"
)

type Location struct {
	Name string
	Lat  float64
	Lng  float64
}

func (a *ArtistDetail) RenderArtistDetailPage() error {
	fs := []string{
		"internal/web/static/pages/artist_profile.html",
	}

	id := a.atoi(chi.URLParam(a.request, "id"))

	artistInfo := a.client.GetByIdKey()[id]

	data := struct {
		HomeUrl, ArtistDetailUrl, AllEventsPageUrl, TicketUrl, PathUrl string
		ArtistInfo                                                     artistapi.ArtistInfo
		AllArtists                                                     map[int]artistapi.ArtistInfo
		JsObject                                                       template.JS
		ArtistIdKey, DateKey, PathKey, LocationKey                     string
	}{
		LocationKey:      utils.LOCATION_KEY,
		PathKey:          utils.PATH_KEY,
		DateKey:          utils.DATE_KEY,
		ArtistIdKey:      utils.ARTIST_ID_KEY,
		PathUrl:          a.request.URL.EscapedPath(),
		TicketUrl:        utils.TICKET.String(),
		HomeUrl:          utils.HOME.String(),
		ArtistInfo:       artistInfo,
		AllArtists:       a.client.GetByIdKey(),
		ArtistDetailUrl:  utils.ARTIST_DETAILS.String(),
		AllEventsPageUrl: utils.ALL_EVENTS_ROUTES.String(),
	}

	temp := template.Must(template.New("artist_profile.html").Funcs(a.detailPageFuncMap()).ParseFS(a.embedded.Get(), fs...))

	if err := temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceR,
		})
	}
	return nil
}
