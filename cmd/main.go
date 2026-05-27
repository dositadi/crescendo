package main

import (
	"context"
	"fmt"

	artistapi "github.com/dositadi/groupie-tracker/cmd/api/artist_api"
)

type Test struct{ Name string }

func main() {
	/* app := &app.App{}
	app.Run() */

	a := artistapi.ArtistInfo{}

	artists, _ := a.FetchArtists()

	ch := make(chan artistapi.ArtistInfo, 52)
	chError := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())

	artistFilled := a.FillArtistsInfoFromArtists(ch, artists)
	locationFilled := a.FillArtistInfoFromLocation(ctx, artistFilled, chError, artists)
	dateFilled := a.FillArtistInfoFromDate(ctx, locationFilled, chError, artists)

	fmt.Println(len(dateFilled))

	for v := range locationFilled {
		fmt.Println("location: ", v)
	}

	for v := range dateFilled {
		fmt.Println("date: ", v)
	}

	select {
	case <-chError:
		cancel()
	default:
		return
	}
}
