package authpost

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/utils"
)

func (a *Auth) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     utils.ACCESS_TOKEN_KEY,
		Value:    "",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, utils.LOGIN.String(), http.StatusSeeOther)
}
