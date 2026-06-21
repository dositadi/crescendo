package getauth

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/authservice"
)

func (a *Auth) SettingsPageHandler(w http.ResponseWriter, r *http.Request) {
	page := authservice.New(w, a.embedded, a.logger)

	if err := page.RenderSettingsPage(); err != nil {
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": "getauth.SettingsPageHandler()",
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
