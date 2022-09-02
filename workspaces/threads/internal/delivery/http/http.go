package api

import (
	get "threads/internal/delivery/http/GET"

	"github.com/go-chi/chi/v5"
)

func Root(router chi.Router) {
	// TODO: to prevent overfetching
	// create specific routes that are meant to update and query specific fields
	//	router.Post(pattern string, h http.HandlerFunc)

	// don't add anything here, all GET/PUT/DELETE/POST routes should go in their respective packages
	router.Route("/", func(router chi.Router) { get.Routes(router) })
	//	router.Put()
	//	router.Delete()
}
