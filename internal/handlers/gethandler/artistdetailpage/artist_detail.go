package artistdetailpage

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/handlers/gethandler/homepage"
	"github.com/dositadi/groupie-tracker/internal/handlers/posthandler/ticketpage"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type DetailPage struct {
	logger          jsonlog.Logger
	usermodel       homepage.UserModel
	client          herokuapp.HerokuApp
	embedded        groupietracker.Embedded
	favoriteModel   homepage.FavoriteModel
	preferencemodel homepage.PreferenceModel
	soldTickets     ticketpage.SoldTicketsModel
}

func New(usermodel homepage.UserModel, client herokuapp.HerokuApp, logger jsonlog.Logger, embedded groupietracker.Embedded, favoriteModel homepage.FavoriteModel, preferencemodel homepage.PreferenceModel, soldTickets ticketpage.SoldTicketsModel) *DetailPage {
	return &DetailPage{
		usermodel:       usermodel,
		client:          client,
		logger:          logger,
		embedded:        embedded,
		favoriteModel:   favoriteModel,
		preferencemodel: preferencemodel,
		soldTickets:     soldTickets,
	}
}
