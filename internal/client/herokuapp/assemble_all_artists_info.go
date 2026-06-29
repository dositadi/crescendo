package herokuapp

import (
	"context"
	"os"
	"time"
)

var (
	byId = make(map[int]ArtistInfo)
)

func (h *HerokuApp) mapArtistsInfo() {
	// Using the pipeline routine pattern to generate the artist's info
	arts, err := h.fetchArtists()
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chError := make(chan error)

	filledArtists := h.fillArtistsInfoFromArtists(arts)
	chArtistInfo := h.fillArtistInfoFromLocation(ctx, filledArtists, chError, arts)
	chArtistInfo = h.fillArtistInfoFromDate(ctx, chArtistInfo, chError, arts)
	chArtistInfo = h.fillArtistInfoFromRelations(ctx, chArtistInfo, chError, arts)
	//chArtistInfo = h.fillGeolocationsFromOpenCage(ctx, chArtistInfo, chError)

	select {
	case <-chError:
		time.Sleep(5 * time.Millisecond)
		os.Exit(1)
	default:
	}
	for artistInfo := range chArtistInfo {
		byId[artistInfo.Id] = *artistInfo
	}
}
