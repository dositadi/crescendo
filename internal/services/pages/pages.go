package pages

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	favoritemodel "github.com/dositadi/groupie-tracker/internal/models/favorite_model"
)

type Filter string
type Sort string
type Favorite string

const (
	// Filters
	FILTER_BY_ID            Filter = "ID"
	FILTER_BY_NAME          Filter = "NAME"
	FILTER_BY_CREATION_DATE Filter = "CREATION DATE"
	FILTER_BY_FIRST_ALBUM   Filter = "FIRST ALBUM"

	// Sort orders
	ASCENDING_ORDER  Sort = "ASC"
	DESCENDING_ORDER Sort = "DESC"

	// Favorite
	FAVORITED     Favorite = "true"
	NOT_FAVORITED Favorite = "false"
)

type Pages struct {
	logger         jsonlog.Logger
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	client         artistapi.ArtistInfo
	request        *http.Request
	favoriteModel  favoritemodel.FavoriteModel
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client artistapi.ArtistInfo, request *http.Request, favoriteModel favoritemodel.FavoriteModel) *Pages {
	return &Pages{
		logger:         logger,
		responseWriter: responseWriter,
		embedded:       embedded,
		client:         client,
		request:        request,
		favoriteModel:  favoriteModel,
	}
}

func (p *Pages) getUserId()

func (p *Pages) updateClient() {
	p.favoriteModel.GetAll()
}
