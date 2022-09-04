package router

import (
	"log"
	"net/http"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"

	"authz"
)

type RouterConfig struct {
	Enforcer       *casbin.Enforcer
	RootRoute      func(router chi.Router)
	Origin         string
	AllowedOrigins []string
	db             *sqlx.DB
}

func (routerConfig *RouterConfig) Set() *chi.Mux {

	if routerConfig.RootRoute == nil {
		log.Fatal("a root route must be provided")
	}
	if routerConfig.Origin == "" {
		log.Fatal("Origin must be provided")
	}

	router := chi.NewRouter()

	//https://go-chi.io/#/pages/middleware
	router.Use(middleware.RealIP)
	router.Use(middleware.CleanPath)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.AllowContentEncoding("gzip", "deflate"))
	allowedCharsets := []string{"UTF-8", "utf-8", ""}
	router.Use(middleware.ContentCharset(allowedCharsets...))
	router.Use(middleware.RouteHeaders().Route("Origin", routerConfig.Origin, cors.Handler(cors.Options{
		AllowedOrigins:   routerConfig.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})).Handler)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Throttle(10))
	router.Use(middleware.Timeout(time.Second * 60))
	router.Use(passDBToHandlersThroughContext(routerConfig.db))
	if routerConfig.Enforcer != nil {
		router.Use(authz.Authorizor(routerConfig.Enforcer))
		// we need to reset log prefix because authz package has its own
		log.SetPrefix("[Accounts Service]: ")
	}
	router.NotFound(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.Header().Set("User-Agent", "social-mesh")
		writer.WriteHeader(404)
		writer.Write([]byte("Route does not exist"))
	})
	router.MethodNotAllowed(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "text/plain")
		writer.Header().Set("User-Agent", "social-mesh")
		writer.WriteHeader(405)
		writer.Write([]byte("Method is not allowed"))
	})

	// this should be the only route registered here
	// don't add random routes here
	router.Route("/", routerConfig.RootRoute)

	return router
}
