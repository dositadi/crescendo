package pages

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceRHome = "Render Home page f(n) under pages pkg"
	PageHome    = "Home"
	PageFav     = "Favorite"
)

var currentPage = 1
var currentIndex = 0
var count = 0

func (p *Pages) RenderHomePage(partial bool) error {
	fs := []string{
		"internal/web/static/pages/home_page.html",
		"internal/web/static/partials/pages/home_page_partials.html",
	}

	userFavorites, err := p.getUserFavorites()
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return err
	}

	userPreference, err := p.getUserPreference()
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return err
	}

	temp, err := template.New("home_page.html").Funcs(p.homePageFunc()).ParseFS(p.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return e
	}

	var artists []herokuapp.ArtistInfo
	var paginatedArtists []herokuapp.ArtistInfo

	artists = sortArtist(mapToSlice(p.client.Get()), Sort(userPreference.Sort), Filter(userPreference.Filter))

	var page string

	if !partial {
		page = p.request.FormValue(utils.PAGE_KEY)
	}

	if page != "" {
		p := p.atoi(page)
		if p < 0 {
			currentPage += p
		} else {
			currentPage = p
		}
	}

	var disableNextButton bool
	var disablePrevButton bool
	artistsLen := len(artists) - 1

	limit := 10
	offset := (currentPage - 1) * limit
	currentIndex = offset + limit

	if currentIndex < artistsLen {
		paginatedArtists = artists[offset : offset+limit]
		count = artistsLen - (artistsLen - (offset + limit))
	} else {
		if offset < artistsLen {
			paginatedArtists = artists[offset:]
			count = artistsLen + 1
			disableNextButton = true
		} else if offset == artistsLen {
			disableNextButton = true
		}
	}

	if currentPage == 1 {
		disablePrevButton = true
	}

	data := struct {
		Username                                                                                   string
		NextPage, PreviousPage, Count, Total                                                       int
		UserFavorites                                                                              map[int]data.Favorite
		Artists                                                                                    []herokuapp.ArtistInfo
		CurrentFilter, CurrentSort                                                                 string
		FilterSortRoute                                                                            string
		FilterByName, FilterByCreationDate, FilterByFirstAlbum                                     string
		FilterKey, ArtistIDKey, SearchKey, PageKey                                                 string
		SortKey, SortASC, SortDESC                                                                 string
		FavoriteArtistUrl, AllFavoritesUrl, FavKey, Favorited, NotFavorited, RequestPage           string
		SearchUrl, Url, ArtistDetailUrl, PrivacyPageUrl, TermsPageUrl, AboutPageUrl, PaidEventsUrl string
		DisableNextbutton, DisablePrevButton, IsSearch                                             bool
	}{
		Username:             strings.Fields(p.getUser().Username)[0],
		UserFavorites:        userFavorites,
		Artists:              paginatedArtists,
		CurrentFilter:        userPreference.Filter,
		CurrentSort:          userPreference.Sort,
		FilterSortRoute:      utils.FILTER_SORT_ROUTE.String(),
		FilterByName:         string(FILTER_BY_NAME),
		FilterByCreationDate: string(FILTER_BY_CREATION_DATE),
		FilterByFirstAlbum:   string(FILTER_BY_FIRST_ALBUM),
		FilterKey:            utils.FILTER_KEY,
		SortKey:              utils.SORT_KEY,
		SortASC:              string(ASCENDING_ORDER),
		SortDESC:             string(DESCENDING_ORDER),
		FavoriteArtistUrl:    utils.FAVORITE.String(),
		AllFavoritesUrl:      utils.FAVORITES.String(),
		FavKey:               utils.FAV_KEY,
		Favorited:            string(FAVORITED),
		NotFavorited:         string(NOT_FAVORITED),
		ArtistIDKey:          utils.ARTIST_ID_KEY,
		SearchUrl:            utils.ARTIST_SEARCH.String(),
		SearchKey:            utils.SEARCH_KEY,
		Url:                  utils.HOME.String(),
		PageKey:              utils.PAGE_KEY,
		NextPage:             currentPage + 1,
		PreviousPage:         currentPage - 1,
		Count:                count,
		Total:                len(artists),
		DisableNextbutton:    disableNextButton,
		DisablePrevButton:    disablePrevButton,
		IsSearch:             false,
		ArtistDetailUrl:      utils.ARTIST_DETAILS.String(),
		PrivacyPageUrl:       fmt.Sprintf("%s?%s=%s", utils.PRIVACY.String(), utils.PAGE_KEY, p.request.URL.EscapedPath()),
		TermsPageUrl:         fmt.Sprintf("%s?%s=%s", utils.TERMS.String(), utils.PAGE_KEY, p.request.URL.EscapedPath()),
		AboutPageUrl:         fmt.Sprintf("%s?%s=%s", utils.ABOUT.String(), utils.PAGE_KEY, p.request.URL.EscapedPath()),
		PaidEventsUrl:        utils.ALL_EVENTS_ROUTES.String(),
		RequestPage:          PageHome,
	}

	p.responseWriter.WriteHeader(http.StatusOK)

	if partial {
		if err = temp.ExecuteTemplate(p.responseWriter, "artist-card-main", data); err != nil {
			e := helper.WrapError("Error executing template", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceRHome,
			})
			return e
		}
	} else {
		if err = temp.Execute(p.responseWriter, data); err != nil {
			e := helper.WrapError("Error executing template", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceRHome,
			})
			return e
		}
	}

	return nil
}

/* func (p *Pages) isHTMXRequest() bool {
	return p.request.Header.Get("HX-Request") == "true"
} */
