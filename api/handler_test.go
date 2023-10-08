package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/yuvrajsingh79/url-shortner/pkg/controller"
)

func TestShortenURLHandler(t *testing.T) {
	// Initialize the URL shortener with mock storage and cache.
	storage := controller.NewMockStorage()
	cache := controller.NewMockCache()
	shortener := controller.NewShortener(storage, cache)

	// Create a test server with the API routes.
	router := SetupRoutes(shortener)
	server := httptest.NewServer(router)
	defer server.Close()

	t.Run("ShortenURL_ValidRequest_ReturnsShortURL", func(t *testing.T) {
		payload := `{"url":"https://example.com"}`
		resp, err := http.Post(server.URL+"/shorten", "application/json", strings.NewReader(payload))
		if err != nil {
			t.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d; got %d", http.StatusOK, resp.StatusCode)
		}

		// Add assertions for the response body here.
	})

	// Add more test cases for different scenarios.

	// Cleanup or reset the mock storage and cache if needed.
}
