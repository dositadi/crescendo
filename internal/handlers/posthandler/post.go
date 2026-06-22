package posthandler

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/handlers/posthandler/authpost"
	"github.com/dositadi/groupie-tracker/internal/handlers/posthandler/homepagepost"
	"github.com/dositadi/groupie-tracker/internal/handlers/posthandler/ticketpage"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type UserModel interface {
	Delete(id string) error
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
	Insert(user data.User) error
	Update(id string, info data.UpdateUser) error
	EmailExists(email string) (bool, error)
	IDExists(id string) (bool, error)
}

type Post struct {
	logger     jsonlog.Logger
	userModel  UserModel
	client     herokuapp.HerokuApp
	embedded   groupietracker.Embedded
	Auth       authpost.Auth
	HomePage   homepagepost.HomePage
	TicketPage ticketpage.TicketPage
}

func New(userModel UserModel, favoriteModel homepagepost.FavoriteModel, preferenceModel homepagepost.PreferenceModel, soldTicketsModel ticketpage.SoldTicketsModel, storageModel authpost.StorageModel, client herokuapp.HerokuApp, logger jsonlog.Logger, embedded groupietracker.Embedded) *Post {
	return &Post{
		userModel:  userModel,
		logger:     logger,
		client:     client,
		embedded:   embedded,
		Auth:       *authpost.New(logger, userModel, preferenceModel, storageModel, embedded),
		HomePage:   *homepagepost.New(logger, userModel, favoriteModel, preferenceModel, embedded, client),
		TicketPage: *ticketpage.New(logger, embedded, client, soldTicketsModel),
	}
}
