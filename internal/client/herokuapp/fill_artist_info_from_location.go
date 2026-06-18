package herokuapp

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

func (h *HerokuApp) fillArtistInfoFromLocation(ctx context.Context, chArtistInfo chan ArtistInfo, chError chan error, artists map[int]artist) chan *ArtistInfo {
	temp := make(chan *ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	if chArtistInfo == nil || artists == nil {
		h.logger.PrintFatal("Recieved a nil paramter [chArtist | artists]", map[string]string{
			"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
		})
		os.Exit(1)
		return nil
	}

	for artInfo := range chArtistInfo {
		art := artists[artInfo.Id]

		wg.Add(1)

		go func(aInfo ArtistInfo, a artist) {
			defer wg.Done()

			locations, err := fetchInfo[location](a.Locations)
			if err != nil {
				e := helper.WrapError("Fetch info error", err)

				h.logger.PrintError(e.Error(), map[string]string{
					"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
				})

				select {
				case chError <- e:
				case <-ctx.Done():
				default:
				}
			}

			if err = ctx.Err(); err != nil {
				e := helper.WrapError("Stopping location fetch worker routine", err)

				h.logger.PrintFatal(e.Error(), map[string]string{
					"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
					"Worker": fmt.Sprintf("Location filler for %v with %v", aInfo, a),
				})
				return
			}

			artInfo := populateArtistInfo(locations, &aInfo)

			select {
			case temp <- artInfo:
			case <-ctx.Done():
			}

		}(artInfo, art)
	}

	go func() {
		wg.Wait()
		close(temp)
	}()

	h.logger.PrintInfo("Filled in locations into artist's info successfully", map[string]string{
		"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
	})

	return temp
}
