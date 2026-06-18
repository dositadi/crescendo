package getauth

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/utils"
)

func (a *Auth) LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, utils.LOGIN.String(), http.StatusSeeOther)
}
