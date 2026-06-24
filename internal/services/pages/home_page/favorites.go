package pages

import (
	"fmt"
	"html/template"
	"strings"

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
		ArtistDetailUrl, PrevPageUrl, FavoriteArtistUrl, PathUrl, SettingsUrl string
		FavKey, ArtistIDKey, ReqPgKey, PathKey                                string
		Artists                                                               []herokuapp.ArtistInfo
		Favorited                                                             string
		RequestPage                                                           string
		User                                                                  data.User
		Username                                                              string
	}{
		ArtistDetailUrl:   utils.ARTIST_DETAILS.String(),
		FavoriteArtistUrl: utils.FAVORITE.String(),
		PrevPageUrl:       utils.HOME.String(),
		SettingsUrl:       fmt.Sprintf("%s?%s=%s", utils.SETTINGS.String(), utils.PATH_KEY, p.request.URL.EscapedPath()),
		PathUrl:           p.request.URL.EscapedPath(),
		Favorited:         string(FAVORITED),

		FavKey:      utils.FAV_KEY,
		ArtistIDKey: utils.ARTIST_ID_KEY,
		ReqPgKey:    utils.REQ_PAGE_KEY,
		PathKey:     utils.PATH_KEY,

		Artists: artists,

		RequestPage: PageFav,

		User:     user,
		Username: strings.Fields(p.getUser().Username)[0],
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
		if artist, ok := artists[favorite.ArtistId]; ok && favorite.Status {
			favoriteArtists = append(favoriteArtists, artist)
		}
	}
	return favoriteArtists
}
