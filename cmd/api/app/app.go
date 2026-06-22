package app

import (
	"os"

	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/client/herokuapp"
	opencage "github.com/dositadi/groupie-tracker/internal/client/open_cage"
	"github.com/dositadi/groupie-tracker/internal/handlers"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/middlewares"
	"github.com/dositadi/groupie-tracker/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type App struct {
	db        *pgx.Conn
	config    config
	logger    *jsonlog.Logger
	handler   *handlers.Handler
	midleware *middlewares.Middleware
	opencage  *opencage.OpenCage
	client    *herokuapp.HerokuApp
	router    *chi.Mux
	embedded  groupietracker.Embedded
}

func (a *App) init() {
	a.embedded = *groupietracker.New()
	a.logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	a.router = chi.NewRouter()
	a.config.Init(*a.logger)
	a.opencage = opencage.New(a.config.opencageKey)
	a.client = herokuapp.New(*a.logger, *a.opencage)
	a.client.Init()
	a.initSupabase(a.config.supabaseUrl, a.config.supabaseSecretKet)
	a.initDB()
	models := models.New(a.db, *a.logger)
	a.handler = handlers.New(*a.logger, &models.UserModel, &models.FavoriteModel, &models.PreferenceModel, &models.SoldTicketsModel, *a.client, a.embedded)
	a.midleware = middlewares.New(*a.handler, *a.logger, a.embedded)
	a.initHandlers()
}

func (a *App) Run() {
	a.init()
	a.startServer()
}
