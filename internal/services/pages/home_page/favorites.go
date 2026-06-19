package pages

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceF = "pages.RenderFavoritePage()"
)

func (p *Pages) RenderFavoritePage() error {
	fs := []string{
		"internal/web/static/pages/favorites.html",
	}

	user := p.getUser()

	favorites, err := p.favoriteModel.GetAll(user.Id)
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceF,
		})
		return err
	}

	artists := allFavorites(favorites, p.client.Get())

	data := struct {
		ArtistDetailUrl, PrevPageUrl, FavoriteArtistUrl string
		FavKey                                          string
		Artists                                         []herokuapp.ArtistInfo
		FavoritedArtists                                []herokuapp.ArtistInfo
	}{
		ArtistDetailUrl:   utils.ARTIST_DETAILS.String(),
		FavoriteArtistUrl: utils.FAVORITE.String(),
		PrevPageUrl:       utils.HOME.String(),

		FavKey: utils.FAV_KEY,

		Artists: artists,
	}

	temp, err := template.New("favorites.html").Funcs(p.homePageFunc()).ParseFS(p.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template parse error", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceF,
		})
		return err
	}

	if err = temp.Execute(p.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceF,
		})
		return err
	}
	return nil
}

func allFavorites(favorites []data.Favorite, artists map[int]herokuapp.ArtistInfo) []herokuapp.ArtistInfo {
	var favoriteArtists []herokuapp.ArtistInfo
	for _, favorite := range favorites {
		if artist, ok := artists[favorite.ArtistId]; ok {
			favoriteArtists = append(favoriteArtists, artist)
		}
	}
	return favoriteArtists
}
