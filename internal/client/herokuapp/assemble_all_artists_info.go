package herokuapp

import (
	"context"
	"fmt"
	"os"
	"sync"
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

	start := time.Now()

	//filledArtists := h.fillArtistsInfoFromArtists(arts)
	//chArtistInfo := h.fillArtistInfoFromLocation(ctx, filledArtists, chError, arts)
	//chArtistInfo = h.fillArtistInfoFromDate(ctx, chArtistInfo, chError, arts)
	//chArtistInfo = h.fillArtistInfoFromRelations(ctx, chArtistInfo, chError, arts)
	chArtistInfo := h.fillGeolocationsFromOpenCage(ctx, h.fillArtistInfoFromRelations(ctx, h.fillArtistInfoFromDate(ctx, h.fillArtistInfoFromLocation(ctx, h.fillArtistsInfoFromArtists(ctx, arts), chError, arts), chError, arts), chError, arts), chError)

	select {
	case <-chError:
		time.Sleep(5 * time.Millisecond)
		os.Exit(1)
	default:
	}
	for artistInfo := range chArtistInfo {
		byId[artistInfo.Id] = *artistInfo
	}
	fmt.Println(time.Since(start))
}

func fanIn(ctx context.Context, workers ...<-chan *ArtistInfo) <-chan *ArtistInfo {
	multiplexedStream := make(chan *ArtistInfo)
	multiplexGroup := new(sync.WaitGroup)

	multiplex := func(chArtist <-chan *ArtistInfo) {
		defer multiplexGroup.Done()
		for artist := range chArtist {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- artist:
			}
		}
	}

	multiplexGroup.Add(len(workers))
	for _, worker := range workers {
		go multiplex(worker)
	}

	go func() {
		multiplexGroup.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func fanOut(chArtists chan *ArtistInfo) []<-chan *ArtistInfo {
	workerSize := 10
	workers := make([]<-chan *ArtistInfo, workerSize)

	// Note: each channel is a product of a go routine and therefore is a go routine itself. Hence, I have 10 worker routines fetching the data
	for i := range workerSize {
		workers[i] = chArtists
	}
	return workers
}
