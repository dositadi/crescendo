package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

const (
	sourceServer = "Start server f(n) under app pkg"
)

func (a *App) startServer() {
	server := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      a.router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		a.logger.PrintInfo("Server running", map[string]string{
			"Source": sourceServer,
			"Port":   "http://localhost:8080/auth/session",
		})
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e := helper.WrapError("Server error", err)
			a.logger.PrintFatal(e.Error(), map[string]string{
				"Source": sourceServer,
			})
			os.Exit(1)
		}
	}()

	<-chSignal

	chError := make(chan error)

	go func() {
		chError <- server.Shutdown(ctx)
	}()

	if err := <-chError; err != nil {
		e := helper.WrapError("Server shutdown forcefully", err)
		a.logger.PrintFatal(e.Error(), map[string]string{
			"Source": sourceServer,
		})
		os.Exit(1)
	}

	a.db.Close(ctx)

	a.logger.PrintInfo("Server shutdown gracefully", map[string]string{
		"Source": sourceServer,
	})
}
