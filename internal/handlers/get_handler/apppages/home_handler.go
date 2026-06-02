package apppages

import (
	"fmt"
	"net/http"
)

func (a *Pages) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page.")
}
