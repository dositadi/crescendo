package herokuapp

import (
	"context"
	"fmt"
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

	start := time.Now()

	chArtistInfo := h.fillArtistsInfoFromArtists(ctx, arts)
	chArtistInfo = h.fillArtistInfoFromDate(ctx, chArtistInfo, chError, arts)
	chArtistInfo = h.fillArtistInfoFromLocation(ctx, chArtistInfo, chError, arts)
	chArtistInfo = h.fillArtistInfoFromRelations(ctx, chArtistInfo, chError, arts)
	//chArtistInfo = h.fillGeolocationsFromOpenCage(ctx, chArtistInfo, chError)

	for artistInfo := range h.orDone(ctx, cancel, chError, chArtistInfo) {
		byId[artistInfo.Id] = *artistInfo
	}

	fmt.Println(time.Since(start))
}

func (h *HerokuApp) orDone(ctx context.Context, cancel context.CancelFunc, done <-chan error, chArtistInfo <-chan *ArtistInfo) <-chan *ArtistInfo {
	out := make(chan *ArtistInfo)

	go func() {
		defer close(out)

		for {
			select {
			case e, ok := <-done:
				if !ok {
					return
				}
				h.logger.PrintError(e.Error(), map[string]string{
					"Source": "herokuapp.mapArtistsInfo()",
				})
				cancel()
				os.Exit(1)
			case artistInfo, ok := <-chArtistInfo:
				if !ok {
					return
				}
				
				select {
				case out <- artistInfo:
				case <-ctx.Done():
				}
			}
		}
	}()
	return out
}

/* func fanIn(ctx context.Context, workers ...<-chan *ArtistInfo) <-chan *ArtistInfo {
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

func fanOut(in <-chan *ArtistInfo, wk func(*ArtistInfo) *ArtistInfo) []<-chan *ArtistInfo {
	workerSize := 10
	outputs := make([]<-chan *ArtistInfo, workerSize)

	// Note: each channel is a product of a go routine and therefore is a go routine itself. Hence, I have 10 worker routines fetching the data
	for i := range workerSize {
		out := make(chan *ArtistInfo)
		outputs[i] = out

		go func(out chan *ArtistInfo) {
			for artist := range in {
				out <- wk(artist)
			}
		}(out)
	}
	return outputs
}
*/
