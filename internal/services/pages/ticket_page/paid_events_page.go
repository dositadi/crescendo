package ticketpage

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
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
		User                                                                            data.User
		Username                                                                        string
		PaidEvents                                                                      []data.SoldTickets
		ArtistInfoName, ArtistInfoImage                                                 string
		PrivacyPageUrl, TermsPageUrl, PrevPageUrl, PathUrl, RecieptPageUrl, SettingsUrl string

		PathKey, DateKey, ArtistIdKey, LocationKey string
	}{
		User:            user,
		Username:        strings.Fields(user.Username)[0],
		PaidEvents:      boughtTickets,
		ArtistInfoName:  "name",
		ArtistInfoImage: "image",
		PrivacyPageUrl:  fmt.Sprintf("%s?%s=%s", utils.PRIVACY.String(), utils.PAGE_KEY, t.request.URL.EscapedPath()),
		TermsPageUrl:    fmt.Sprintf("%s?%s=%s", utils.TERMS.String(), utils.PAGE_KEY, t.request.URL.EscapedPath()),
		PrevPageUrl:     utils.HOME.String(),
		PathUrl:         t.request.URL.EscapedPath(),
		RecieptPageUrl:  utils.RECIEPTS.String(),
		SettingsUrl:     fmt.Sprintf("%s?%s=%s", utils.SETTINGS.String(), utils.PATH_KEY, t.request.URL.EscapedPath()),

		PathKey:     utils.PATH_KEY,
		DateKey:     utils.DATE_KEY,
		ArtistIdKey: utils.ARTIST_ID_KEY,
		LocationKey: utils.LOCATION_KEY,
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
