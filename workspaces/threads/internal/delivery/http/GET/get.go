package get

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(router chi.Router) {
	// all get routes are belong to this function

	router.Get("/", func(rw http.ResponseWriter, req *http.Request) {

		//getThreadByID(rw, req)
	})
}
