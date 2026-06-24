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
	sourceE = "Render all events page f(n) under artistdetail pkg"
)

func (a *ArtistDetail) RenderAllEventsPage() error {
	fs := []string{
		"internal/web/static/pages/all_concerts.html",
	}

	id := a.atoi(chi.URLParam(a.request, "id"))
	user := a.getUser()

	artistInfo := a.client.Get()[id]

	userTickets, err := a.soldTicketsModel.GetAll(user.Id)
	if err != nil {
		e := helper.WrapError("User tickets get err", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceR,
		})
	}

	prevPage := a.request.URL.EscapedPath()

	detailPage := fmt.Sprintf("%s/%v", utils.ARTIST_DETAILS.String(), id)

	data := struct {
		ArtistDetailUrl, TicketUrl, PathUrl, RecieptPageUrl, DetailPageUrl,SettingsUrl string
		ArtistInfo                                                         herokuapp.ArtistInfo
		AllArtists                                                         map[int]herokuapp.ArtistInfo
		UserTickets                                                        []data.SoldTickets
		PathKey, DateKey, ArtistIdKey, LocationKey                         string

		User data.User
		Username string
	}{
		User:     user,
		Username: strings.Fields(a.getUser().Username)[0],

		// All urls
		PathUrl:         prevPage,
		DetailPageUrl:   detailPage,
		ArtistDetailUrl: utils.ARTIST_DETAILS.String(),
		TicketUrl:       utils.TICKET.String(),
		RecieptPageUrl:  utils.RECIEPTS.String(),
		SettingsUrl:     fmt.Sprintf("%s?%s=%s", utils.SETTINGS.String(), utils.PATH_KEY, a.request.URL.EscapedPath()),

		// Artists details
		ArtistInfo:  artistInfo,
		AllArtists:  a.client.Get(),
		UserTickets: userTickets,

		// All keys
		PathKey:     utils.PATH_KEY,
		DateKey:     utils.DATE_KEY,
		ArtistIdKey: utils.ARTIST_ID_KEY,
		LocationKey: utils.LOCATION_KEY,
	}

	temp, err := template.New("all_concerts.html").Funcs(a.detailPageFuncMap()).ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template create error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Status": sourceE,
		})
	}

	if err = temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Status": sourceE,
		})
	}
	return nil
}
