package pages

import (
	"net/http"
	"strconv"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/pages"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/google/uuid"
)

const (
	sourceUH = "Update favorite handler under pages pkg"
)

func (p *Pages) UpdateFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	favStatus := r.FormValue(utils.FAV_KEY)
	val := r.FormValue(utils.ARTIST_ID_KEY)
	artistId, _ := strconv.Atoi(val) // Make this a f(n)
	idVal := r.Context().Value(utils.USER_ID_KEY)
	var userId = ""

	if id, ok := idVal.(string); ok {
		userId = id
	}

	switch favStatus {
	case string(pages.FAVORITED):
		exists, err := p.favoriteModel.Exists(artistId)
		if err != nil {
			e := helper.WrapError("Update favorite error", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUH,
			})
			return
		}

		id := uuid.NewString()
		favorite := data.Favorite{
			Id:       id,
			UserId:   userId,
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
			status := false
			favUpdate := data.FavoriteUpdate{
				UserId:   userId,
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
		exists, err := p.favoriteModel.Exists(artistId)
		if err != nil {
			e := helper.WrapError("Update favorite error", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUH,
			})
			return
		}

		id := uuid.NewString()
		favorite := data.Favorite{
			Id:       id,
			UserId:   userId,
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
			status := false
			favUpdate := data.FavoriteUpdate{
				UserId:   userId,
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
	
}
