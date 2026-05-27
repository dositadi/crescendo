package app

import (
	"os"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/handlers"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/middleware"
	"github.com/jackc/pgx/v5"
)

type App struct {
	db        *pgx.Conn
	config    config
	logger    *jsonlog.Logger
	handler   *handlers.Handler
	midleware *middleware.Middleware
	client    *artistapi.ArtistInfo
}

func (a *App) init() {
	a.client.Init()
	a.config = newConfig()
	a.logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	a.handler = handlers.New(*a.logger)
	a.midleware = middleware.New(*a.handler, *a.logger)
	a.initDB()
}

func (a *App) Run() {
	a.init()
}
