package middlewares

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/handlers"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type UserModel interface {
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
}

type Middleware struct {
	handler   handlers.Handler
	logger    jsonlog.Logger
	embedded  groupietracker.Embedded
	usermodel UserModel
}

func New(handler handlers.Handler, logger jsonlog.Logger, embedded groupietracker.Embedded, usermodel UserModel) *Middleware {
	return &Middleware{
		handler:   handler,
		logger:    logger,
		embedded:  embedded,
		usermodel: usermodel,
	}
}
