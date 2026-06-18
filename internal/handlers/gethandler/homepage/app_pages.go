package homepage

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/data"
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

type FavoriteModel interface {
	DeleteAll(userId string) error
	Delete(userId string, artistId string) error
	Exists(artistId int, userId string) (bool, error)
	Get(artistId int, userId string) (data.Favorite, error)
	GetAll(userId string) ([]data.Favorite, error)
	Insert(favorite data.Favorite) error
	Update(fav data.FavoriteUpdate) error
}

type PreferenceModel interface {
	Exists(userId string) (bool, error)
	Get(userId string) (data.Preference, error)
	Insert(preference data.Preference) error
	Update(preference data.PreferenceUpdate) error
}

type HomePage struct {
	logger          jsonlog.Logger
	usermodel       UserModel
	client          herokuapp.HerokuApp
	embedded        groupietracker.Embedded
	favoriteModel   FavoriteModel
	preferencemodel PreferenceModel
}

func New(usermodel UserModel, client herokuapp.HerokuApp, logger jsonlog.Logger, embedded groupietracker.Embedded, favoriteModel FavoriteModel, preferencemodel PreferenceModel) *HomePage {
	return &HomePage{
		usermodel:       usermodel,
		client:          client,
		logger:          logger,
		embedded:        embedded,
		favoriteModel:   favoriteModel,
		preferencemodel: preferencemodel,
	}
}
