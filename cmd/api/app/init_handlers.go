package app

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) fileServer() {
	a.router.Handle("/src/output.css", http.FileServerFS(a.embedded.Get()))
}

func (a *App) initHandlers() {
	a.fileServer()
	a.router.Use(middleware.CleanPath)

	// Get request routes
	a.router.Group(func(r chi.Router) {
		r.With(a.midleware.VerifyAccessToken).Get(utils.LOGIN.String(), a.handler.Get.Auth.LoginHandler)
	})
}
