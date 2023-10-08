package controller

import (
	"errors"
	"fmt"
	"net/url"
	"sync"

	"github.com/google/uuid"
)

// BaseURL is the base URL for your short URLs.
const BaseURL = "https://example.com" // Replace with your actual base URL

// Shortener is a service that generates and retrieves shortened URLs.
type Shortener struct {
	storage Storage    // Use the Storage interface for storing and retrieving URLs.
	cache   Cache      // Use a caching mechanism (e.g., Redis) for performance.
	mu      sync.Mutex // Mutex for ensuring thread safety.
}

// NewShortener creates a new Shortener instance.
func NewShortener(storage Storage, cache Cache) *Shortener {
	return &Shortener{
		storage: storage,
		cache:   cache,
	}
}

// Shorten generates a short URL for the given original URL and stores it using the storage.
func (s *Shortener) Shorten(originalURL string) (string, error) {
	// Validate the URL.
	parsedURL, err := url.Parse(originalURL)
	if err != nil || !parsedURL.IsAbs() {
		return "", errors.New("invalid URL")
	}

	// Check if the URL is already shortened.
	shortURL, err := s.storage.RetrieveShortURL(originalURL)
	if err == nil {
		return shortURL, nil
	}

	// Generate a unique short URL key.
	shortURLKey, err := generateUniqueShortURLKey()
	if err != nil {
		return "", err
	}

	// Create the short URL.
	shortURL = fmt.Sprintf("%s/%s", BaseURL, shortURLKey)

	// Store the mapping in storage.
	if err := s.storage.StoreURLMapping(shortURLKey, originalURL); err != nil {
		return "", err
	}

	// Update the cache with the new mapping.
	s.cache.Set(shortURLKey, originalURL)

	return shortURL, nil
}

// Retrieve retrieves the original URL associated with the provided short URL using the cache and storage.
func (s *Shortener) Retrieve(shortURL string) (string, error) {
	// Check the cache first.
	originalURL, found := s.cache.Get(shortURL)
	if found {
		return originalURL, nil
	}

	// If not in cache, retrieve from storage.
	originalURL, err := s.storage.RetrieveOriginalURL(shortURL)
	if err != nil {
		return "", err
	}

	// Store in cache for future use.
	s.cache.Set(shortURL, originalURL)

	return originalURL, nil
}

// generateUniqueShortURLKey generates a unique short URL key using UUID.
func generateUniqueShortURLKey() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
