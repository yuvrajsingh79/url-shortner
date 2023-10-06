package api

import (
	"interview/url-shortner/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes configures API routes.
func SetupRoutes(s *controller.Shortener) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", ShortenURLHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/{shortURL}", RedirectHandler(s)).Methods(http.MethodGet)

	return r
}
