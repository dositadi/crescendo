package auth
import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/authservice"
)

func (h *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	authService := authservice.New(w, h.embedded, h.logger)

	if err := authService.RenderLoginPage(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
