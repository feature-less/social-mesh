package router

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func passDBToHandlersThroughContext(db *sqlx.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(context.Background(), "db", db)
			next.ServeHTTP(writer, req.WithContext(ctx))
		})
	}
}
