package service

import (
	"net/http"
	"service02/tracing"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(loadTracing(route.HandlerFunc))

	}
	return router
}

func loadTracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		span := tracing.StartHTTPTrace(req, "Hello")
		defer span.Finish()

		ctx := tracing.UpdateContext(req.Context(), span)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
