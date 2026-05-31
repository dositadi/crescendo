package main

import (
	"github.com/dositadi/groupie-tracker/cmd/api/app"
)

func main() {
	app := &app.App{}
	app.Run()
}
