package controller

import (
	"errors"
	"sync"
)

// Storage is an interface representing a data storage mechanism for URL mappings.
type Storage interface {
	Save(shortURL, originalURL string) error
	Retrieve(shortURL string) (string, error)
}

// InMemoryStorage is an in-memory implementation of the Storage interface.
type InMemoryStorage struct {
	urlMap map[string]string // A map to store URL mappings (short URL to original URL).
	mu     sync.RWMutex      // Mutex for concurrent access to the map.
}

// NewInMemoryStorage creates a new InMemoryStorage instance.
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urlMap: make(map[string]string),
	}
}

// Save stores the given short URL and its corresponding original URL.
func (s *InMemoryStorage) Save(shortURL, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urlMap[shortURL] = originalURL
	return nil
}

// Retrieve retrieves the original URL associated with the provided short URL.
func (s *InMemoryStorage) Retrieve(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	originalURL, exists := s.urlMap[shortURL]
	if !exists {
		return "", errors.New("short URL not found")
	}
	return originalURL, nil
}
