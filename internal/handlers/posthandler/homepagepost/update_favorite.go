package homepagepost

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	pages "github.com/dositadi/groupie-tracker/internal/services/pages/home_page"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/google/uuid"
)

const (
	sourceUH = "Update favorite handler under pages pkg"
)

func (p *HomePage) UpdateFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	favStatus := r.FormValue(utils.FAV_KEY)
	val := r.FormValue(utils.ARTIST_ID_KEY)
	artistId := p.atoi(val)
	requestPage := r.FormValue(utils.REQ_PAGE_KEY)
	page := pages.New(p.logger, w, p.embedded, p.client, r, p.favoriteModel, p.preferenceModel)
	user := p.getUserId(r)

	exists, err := p.favoriteModel.Exists(artistId, user.Id)
	if err != nil {
		e := helper.WrapError("Update favorite error", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUH,
		})
		return
	}

	switch favStatus {
	case string(pages.FAVORITED):
		id := uuid.NewString()
		favorite := data.Favorite{
			Id:       id,
			UserId:   user.Id,
			ArtistId: artistId,
			Status:   false,
		}

		switch exists {
		case false:
			if err := p.favoriteModel.Insert(favorite); err != nil {
				e := helper.WrapError("Update favorite (insert) error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		case true:
			fmt.Println("Entered true Fav")
			status := false
			favUpdate := data.FavoriteUpdate{
				UserId:   user.Id,
				ArtistId: artistId,
				Status:   &status,
			}
			if err := p.favoriteModel.Update(favUpdate); err != nil {
				e := helper.WrapError("Update favorite error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		}

	case string(pages.NOT_FAVORITED):
		id := uuid.NewString()
		favorite := data.Favorite{
			Id:       id,
			UserId:   user.Id,
			ArtistId: artistId,
			Status:   true,
		}

		switch exists {
		case false:
			if err := p.favoriteModel.Insert(favorite); err != nil {
				e := helper.WrapError("Update favorite (insert) error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		case true:
			var status bool = true
			favUpdate := data.FavoriteUpdate{
				UserId:   user.Id,
				ArtistId: artistId,
				Status:   &status,
			}
			if err := p.favoriteModel.Update(favUpdate); err != nil {
				e := helper.WrapError("Update favorite error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		}
	}

	switch requestPage {
	case pages.PageHome:
		if err := page.RenderHomePage(true); err != nil {
			e := helper.WrapError("Render favorite button error", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUH,
			})
			return
		}
	case pages.PageFav:
		http.Redirect(w, r, utils.FAVORITES.String()+"#artists-grid", http.StatusSeeOther)
	}
}

func (p *HomePage) atoi(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		p.logger.PrintError("Atoi conversion error: Not a valid number", map[string]string{
			"Source": sourceUH,
		})
		panic("Not a valid number")
	}
	return out
}
