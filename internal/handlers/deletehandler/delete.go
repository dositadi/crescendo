package deletehandler

import (
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

type Delete struct {
	logger    jsonlog.Logger
	userModel UserModel
	client    herokuapp.HerokuApp
}

func New(userModel UserModel, client herokuapp.HerokuApp, logger jsonlog.Logger) *Delete {
	return &Delete{
		userModel: userModel,
		client:    client,
		logger:    logger,
	}
}
