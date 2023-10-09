package controller

import (
	"errors"
)

// Storage defines an interface for storing and retrieving URL mappings.
type Storage interface {
	// StoreURLMapping stores a mapping between a short URL key and an original URL.
	StoreURLMapping(shortURLKey, originalURL string) error

	// RetrieveOriginalURL retrieves the original URL associated with a short URL key.
	RetrieveOriginalURL(shortURLKey string) (string, error)

	// RetrieveShortURL retrieves the short URL associated with an original URL.
	RetrieveShortURL(originalURL string) (string, error)
}

// MemoryStorage is an in-memory implementation of the Storage interface.
type MemoryStorage struct {
	urlMappings map[string]string // Map to store URL mappings.
}

// NewMemoryStorage creates a new MemoryStorage instance.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		urlMappings: make(map[string]string),
	}
}

// StoreURLMapping stores a mapping between a short URL key and an original URL.
func (s *MemoryStorage) StoreURLMapping(shortURLKey, originalURL string) error {
	if _, exists := s.urlMappings[shortURLKey]; exists {
		return errors.New("short URL key already exists")
	}
	s.urlMappings[shortURLKey] = originalURL
	return nil
}

// RetrieveOriginalURL retrieves the original URL associated with a short URL key.
func (s *MemoryStorage) RetrieveOriginalURL(shortURLKey string) (string, error) {
	originalURL, exists := s.urlMappings[shortURLKey]
	if !exists {
		return "", errors.New("short URL key not found")
	}
	return originalURL, nil
}

// RetrieveShortURL retrieves the short URL associated with an original URL.
func (s *MemoryStorage) RetrieveShortURL(originalURL string) (string, error) {
	for key, url := range s.urlMappings {
		if url == originalURL {
			return key, nil
		}
	}
	return "", errors.New("original URL not found")
}
