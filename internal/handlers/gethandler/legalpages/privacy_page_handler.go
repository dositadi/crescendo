package legalpages

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/pages/legal"
)

func (l *LegalPage) PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	page := legal.New(l.logger, w, l.embedded, r)

	if err := page.RenderLegalPages(legal.Privacy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		l.logger.PrintError(err.Error(), map[string]string{
			"Source": "legalpages.PrivacyPageHandler()",
		})
	}
}
