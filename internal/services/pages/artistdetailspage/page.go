package artistdetail

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type ArtistDetail struct {
	logger         jsonlog.Logger
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	client         herokuapp.HerokuApp
	request        *http.Request
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client herokuapp.HerokuApp, request *http.Request) *ArtistDetail {
	return &ArtistDetail{
		logger:         logger,
		responseWriter: responseWriter,
		embedded:       embedded,
		client:         client,
		request:        request,
	}
}
