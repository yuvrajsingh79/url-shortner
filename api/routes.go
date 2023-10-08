package api

import (
	"net/http"

	"github.com/yuvrajsingh79/url-shortner/pkg/controller"

	"github.com/gorilla/mux"
)

// SetupRoutes configures API routes.
func SetupRoutes(s *controller.Shortener) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", ShortenURLHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/{shortURL}", RedirectHandler(s)).Methods(http.MethodGet)

	return r
}
