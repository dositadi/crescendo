package getauth

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

type Auth struct {
	logger    jsonlog.Logger
	usermodel UserModel
	client    herokuapp.HerokuApp
	embedded  groupietracker.Embedded
}

func New(usermodel UserModel, client herokuapp.HerokuApp, logger jsonlog.Logger, embedded groupietracker.Embedded) *Auth {
	return &Auth{
		usermodel: usermodel,
		client:    client,
		logger:    logger,
		embedded:  embedded,
	}
}
