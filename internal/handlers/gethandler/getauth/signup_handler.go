package getauth

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/authservice"
)

func (a *Auth) SignupHandler(w http.ResponseWriter, r *http.Request) {
	authService := authservice.New(w, a.embedded, a.logger)

	if err := authService.RenderSignupPage(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
