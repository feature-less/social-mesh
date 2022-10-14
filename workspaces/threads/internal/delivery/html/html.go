package templates

import (
	_ "threads/docs"

	"github.com/go-chi/chi/v5"

	"threads/internal/delivery/html/controller"
)

// @title Social-Mesh Threads html templates
// @version 1.0
// @description This is an early development server.

// @contact.name Oussama M. Bouchareb
// @contact.email commensalism@proton.me

// @license.name AGPL 3.0
// @host localhost:3000
// @BasePath /

func Root(router chi.Router) {
	// TODO: to prevent overfetching
	// create specific routes that are meant to update and query specific fields
	//	router.Post(pattern string, h http.HandlerFunc)

	// don't add anything here, all GET/PUT/DELETE/POST routes should go in their respective packages
	router.Route("/", func(router chi.Router) { controller.Routes(router) })
}
