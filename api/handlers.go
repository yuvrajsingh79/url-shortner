package api

import (
	"encoding/json"
	"interview/url-shortner/controller"
	"net/http"
	"strings"
)

// ShortenURLHandler handles the shortening of URLs.
func ShortenURLHandler(s *controller.Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			URL string `json:"url"`
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, "Invalid JSON request", http.StatusBadRequest)
			return
		}

		originalURL := strings.TrimSpace(request.URL)
		if originalURL == "" {
			http.Error(w, "URL cannot be empty", http.StatusBadRequest)
			return
		}

		shortURL, err := s.Shorten(originalURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := struct {
			ShortURL string `json:"short_url"`
		}{ShortURL: shortURL}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// RedirectHandler redirects to the original URL.
func RedirectHandler(s *controller.Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		shortURL := r.URL.Path[1:] // Remove the leading '/'
		if shortURL == "" {
			http.Error(w, "Short URL cannot be empty", http.StatusBadRequest)
			return
		}

		originalURL, err := s.Retrieve(shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalURL, http.StatusSeeOther)
	}
}
