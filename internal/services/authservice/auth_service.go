package authservice

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type AuthService struct {
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	logger         jsonlog.Logger
}

func New(w http.ResponseWriter, f groupietracker.Embedded, logger jsonlog.Logger) *AuthService {
	return &AuthService{
		responseWriter: w,
		embedded:       f,
		logger:         logger,
	}
}
