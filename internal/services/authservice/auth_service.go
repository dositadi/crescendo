package authservice

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

type AuthService struct {
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	logger         jsonlog.Logger
	request        *http.Request
}

func New(w http.ResponseWriter, f groupietracker.Embedded, logger jsonlog.Logger, request *http.Request) *AuthService {
	return &AuthService{
		responseWriter: w,
		embedded:       f,
		logger:         logger,
		request:        request,
	}
}

func (a *AuthService) getUser() data.User {
	val := a.request.Context().Value(utils.USER_ID_KEY)
	if user, ok := val.(data.User); ok {
		return user
	}
	return data.User{}
}
