package main

import (
	"os"

	"github.com/dositadi/groupie-tracker/cmd/api/app"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/handlers/post_handler/auth"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

func main() {
	activeUser := data.ActiveUser{
		Id:       "1",
		Username: "Divine",
		Email:    "akindivine587@gmail.com",
	}

	a := auth.New(*jsonlog.New(os.Stdout, jsonlog.LevelInfo), nil)

	a.GenerateJWT(activeUser)
	app := &app.App{}
	app.Run()
}
