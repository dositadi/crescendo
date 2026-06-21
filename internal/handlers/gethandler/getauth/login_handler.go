package getauth

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/authservice"
)

func (h *Auth) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	authService := authservice.New(w, h.embedded, h.logger,r)

	if err := authService.RenderLoginPage(); err != nil {
		h.logger.PrintError(err.Error(), map[string]string{
			"Source": "getauth.LoginPageHandler()",
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
