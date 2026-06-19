package artistdetail

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type SoldTicketsModel interface {
	Exists(userId, date, location string, artistId int) (bool, error)
	Get(artistId int, userId, location, date string) (data.SoldTickets, error)
	GetAll(userId string) ([]data.SoldTickets, error)
	Insert(soldTicket data.SoldTickets) error
}

type ArtistDetail struct {
	logger           jsonlog.Logger
	responseWriter   http.ResponseWriter
	embedded         groupietracker.Embedded
	client           herokuapp.HerokuApp
	request          *http.Request
	soldTicketsModel SoldTicketsModel
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client herokuapp.HerokuApp, request *http.Request, soldTicketsModel SoldTicketsModel) *ArtistDetail {
	return &ArtistDetail{
		logger:           logger,
		responseWriter:   responseWriter,
		embedded:         embedded,
		client:           client,
		request:          request,
		soldTicketsModel: soldTicketsModel,
	}
}
