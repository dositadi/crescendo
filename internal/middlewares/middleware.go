package middlewares

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/handlers"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type Middleware struct {
	handler  handlers.Handler
	logger   jsonlog.Logger
	embedded groupietracker.Embedded
}

func New(handler handlers.Handler, logger jsonlog.Logger, embedded groupietracker.Embedded) *Middleware {
	return &Middleware{
		handler:  handler,
		logger:   logger,
		embedded: embedded,
	}
}
