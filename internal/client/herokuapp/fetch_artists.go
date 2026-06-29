package herokuapp

import (
	"encoding/json"
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

const (
	artistUrl = "https://groupietrackers.herokuapp.com/api/artists"
)

//var logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)

// A generic function that fetches all the artist resource.
func (h *HerokuApp) fetchArtists() (map[int]artist, error) {
	response, err := http.Get(artistUrl)
	if err != nil {
		e := helper.WrapError("Get error", err)
		h.logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch artists f(n) under artistapi package.",
		})
		return nil, e
	}

	defer response.Body.Close()

	var artists []artist

	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		e := helper.WrapError("JSON decode error", err)
		h.logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch artists f(n) under artistapi package.",
		})
		return nil, e
	}

	artistsMap := make(map[int]artist)

	for _, artist := range artists {
		artistsMap[artist.Id] = artist
	}

	h.logger.PrintInfo("Artists fetch successful", map[string]string{
		"Source": "Fetch artists f(n) under artistapi package.",
	})

	return artistsMap, nil
}
