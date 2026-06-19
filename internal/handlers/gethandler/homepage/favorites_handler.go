package homepage

import (
	"net/http"

	pages "github.com/dositadi/groupie-tracker/internal/services/pages/home_page"
)

const (
	sourceF = "homepage.FavoritePageHandler()"
)

func (h *HomePage) FavoritePageHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.New(h.logger, w, h.embedded, h.client, r, h.favoriteModel, h.preferencemodel)

	if err := page.RenderFavoritePage(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceF,
		})
	}
}
