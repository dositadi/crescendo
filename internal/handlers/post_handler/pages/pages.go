package pages

import (
	groupietracker "github.com/dositadi/groupie-tracker"
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
	Exists(artistId int) (bool, error)
	Get(artistId int, userId string)
	GetAll(userId string) ([]data.Favorite, error)
	Insert(favorite data.Favorite) error
	Update(fav data.FavoriteUpdate) error
}

type Pages struct {
	logger        jsonlog.Logger
	usermodel     UserModel
	favoriteModel FavoriteModel
	embedded      groupietracker.Embedded
}

func New(logger jsonlog.Logger, userModel UserModel, favoriteModel FavoriteModel, embedded groupietracker.Embedded) *Pages {
	return &Pages{
		logger:    logger,
		usermodel: userModel,
		embedded:  embedded,
	}
}
