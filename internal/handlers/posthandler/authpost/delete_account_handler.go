package authpost

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

func (a *Auth) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	user := a.getUser(r)

	if err := a.usermodel.Delete(user.Id); err != nil {
		e := helper.WrapError("Delete user error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": "authpost.DeleteAccount()",
		})
		http.Error(w, "Something wrong happened. Kindly try again.", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     utils.ACCESS_TOKEN_KEY,
		Value:    "",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, utils.REGISTER.String(), http.StatusSeeOther)
}
