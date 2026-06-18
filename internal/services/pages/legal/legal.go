package legal

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type LegalPages struct {
	logger         jsonlog.Logger
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	request        *http.Request
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client herokuapp.HerokuApp, request *http.Request) *LegalPages {
	return &LegalPages{
		logger:         logger,
		responseWriter: responseWriter,
		embedded:       embedded,
		request:        request,
	}
}
