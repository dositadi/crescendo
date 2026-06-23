package getauth

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/authservice"
)

func (a *Auth) SignupHandler(w http.ResponseWriter, r *http.Request) {
	authService := authservice.New(w, a.embedded, a.logger, r, a.usermodel)

	if err := authService.RenderSignupPage(r.URL.EscapedPath()); err != nil {
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": "getauth.SignupHandler()",
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
