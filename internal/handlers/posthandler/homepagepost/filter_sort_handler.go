package homepagepost

import (
	"fmt"
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	pages "github.com/dositadi/groupie-tracker/internal/services/pages/home_page"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/google/uuid"
)

const (
	sourceFS = "Filter-Sort handler f(n) under apppages pkg"
)

func (p *HomePage) FilterSortHandler(w http.ResponseWriter, r *http.Request) {
	filter := r.FormValue(utils.FILTER_KEY)
	sort := r.FormValue(utils.SORT_KEY)
	fmt.Println("handler: ", filter, " ", sort)

	user := p.getUserId(r)

	exists, err := p.preferenceModel.Exists(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceFS,
		})
	}

	switch exists {
	case true:
		prefUpdate := data.PreferenceUpdate{
			UserId: user.Id,
			Filter: &filter,
			Sort:   &sort,
		}

		err = p.preferenceModel.Update(prefUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			p.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceFS,
			})
		}
	case false:
		pref := data.Preference{
			Id:     uuid.NewString(),
			UserId: user.Id,
			Filter: filter,
			Sort:   sort,
		}

		err := p.preferenceModel.Insert(pref)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			p.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceFS,
			})
		}
	}

	page := pages.New(p.logger, w, p.embedded, p.client, r, p.favoriteModel, p.preferenceModel)

	if err := page.RenderHomePage(true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceFS,
		})
	}
}
