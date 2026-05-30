package middlewares

import (
	"net/http"
	"os"

	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

var logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)

func (m *Middleware) Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				e, ok := err.(string)
				if !ok {
					e = "An internal server error occurred."
				}
				logger.PrintError(e, map[string]string{
					"Source":     "Recover middleware",
					"PanicRoute": r.URL.Path,
				})
				http.Error(w, e, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
