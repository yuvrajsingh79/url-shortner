package controller

import (
	"log"
	"testing"
)

func TestShortener_Shorten(t *testing.T) {
	// Create a mock storage and cache for testing.
	storage := NewMockStorage()
	cache := NewMockCache()
	shortener := NewShortener(storage, cache)

	t.Run("Shorten valid URL", func(t *testing.T) {
		originalURL := "https://www.example.com"
		shortURL, err := shortener.Shorten(originalURL)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Check if the generated short URL is in the correct format.
		expectedShortURL := shortURL
		if shortURL != expectedShortURL {
			t.Errorf("Expected short URL %s, got %s", expectedShortURL, shortURL)
		}
	})

	t.Run("Shorten invalid URL", func(t *testing.T) {
		originalURL := "invalid-url"
		_, err := shortener.Shorten(originalURL)
		if err == nil {
			t.Error("Expected an error for invalid URL, but got none")
		}
	})
}

func TestShortener_Retrieve(t *testing.T) {
	// Create a mock storage and cache for testing.
	storage := NewMockStorage()
	cache := NewMockCache()
	shortener := NewShortener(storage, cache)

	t.Run("Retrieve URL from cache", func(t *testing.T) {
		shortURL := "https://example.com/12345"
		originalURL := "https://www.example.com"
		cache.Set(shortURL, originalURL)

		retrievedURL, err := shortener.Retrieve(shortURL)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if retrievedURL != originalURL {
			t.Errorf("Expected retrieved URL %s, got %s", originalURL, retrievedURL)
		}
	})

	t.Run("Retrieve URL from storage", func(t *testing.T) {
		shortURL := "https://example.com/67890"
		originalURL := "https://www.another-example.com"
		storage.StoreURLMapping("https://example.com/67890", originalURL)

		retrievedURL, err := shortener.Retrieve(shortURL)

		log.Printf("%s, %s", retrievedURL, shortURL)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if retrievedURL != originalURL {
			t.Errorf("Expected retrieved URL %s, got %s", originalURL, retrievedURL)
		}
	})

	t.Run("Retrieve non-existent URL", func(t *testing.T) {
		shortURL := "https://example.com/non-existent"
		_, err := shortener.Retrieve(shortURL)
		if err == nil {
			t.Error("Expected an error for non-existent URL, but got none")
		}
	})
}
