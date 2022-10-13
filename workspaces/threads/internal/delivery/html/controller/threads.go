package controller

import (
	"fmt"
	"goview"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func examplePage(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "example.html", goview.M{"title": "test"})
	if err != nil {
		fmt.Fprintf(w, "Render index error: %v!", err)
	}
}
func Routes(router chi.Router) {
	// all get routes are belong to this function

	router.Get("/", examplePage)
}
