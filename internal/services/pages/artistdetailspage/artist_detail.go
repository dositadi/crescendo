package artistdetail

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/data"
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
	path := a.request.FormValue(utils.PATH_KEY)

	if path == "" {
		path = utils.HOME.String()
	}

	user := a.getUser()

	artistInfo := a.client.Get()[id]

	userTickets, err := a.soldTicketsModel.GetAll(user.Id)
	if err != nil {
		e := helper.WrapError("User tickets get err", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceR,
		})
	}

	data := struct {
		PrevPageUrl, ArtistDetailUrl, AllEventsPageUrl, TicketUrl, PathUrl, PrivacyPageUrl, TermsPageUrl, RecieptPageUrl, SettingsUrl string
		ArtistInfo                                                                                                                    herokuapp.ArtistInfo
		User                                                                                                                          data.User
		Username                                                                                                                      string
		AllArtists                                                                                                                    map[int]herokuapp.ArtistInfo
		UserTickets                                                                                                                   []data.SoldTickets
		JsObject                                                                                                                      template.JS
		ArtistIdKey, DateKey, PathKey, LocationKey                                                                                    string
	}{
		User:             user,
		Username:         strings.Fields(a.getUser().Username)[0],
		LocationKey:      utils.LOCATION_KEY,
		PathKey:          utils.PATH_KEY,
		DateKey:          utils.DATE_KEY,
		ArtistIdKey:      utils.ARTIST_ID_KEY,
		PathUrl:          a.request.URL.EscapedPath(),
		TicketUrl:        utils.TICKET.String(),
		PrevPageUrl:      path,
		RecieptPageUrl:   utils.RECIEPTS.String(),
		SettingsUrl:      fmt.Sprintf("%s?%s=%s", utils.SETTINGS.String(), utils.PATH_KEY, a.request.URL.EscapedPath()),
		ArtistInfo:       artistInfo,
		AllArtists:       a.client.Get(),
		UserTickets:      userTickets,
		ArtistDetailUrl:  utils.ARTIST_DETAILS.String(),
		AllEventsPageUrl: utils.ALL_EVENTS_ROUTES.String(),
		PrivacyPageUrl:   fmt.Sprintf("%s?%s=%s", utils.PRIVACY.String(), utils.PAGE_KEY, a.request.URL.EscapedPath()),
		TermsPageUrl:     fmt.Sprintf("%s?%s=%s", utils.TERMS.String(), utils.PAGE_KEY, a.request.URL.EscapedPath()),
	}

	temp, err := template.New("artist_profile.html").Funcs(a.detailPageFuncMap()).ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template new error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceR,
		})
	}

	if err = temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceR,
		})
	}
	return nil
}
