package controller

import (
	"errors"
	"sync"
)

// ErrMappingNotFound is an error returned when a mapping is not found.
var ErrMappingNotFound = errors.New("mapping not found")

// ErrCacheMiss is an error returned when a cache miss occurs.
// var ErrCacheMiss = errors.New("cache miss")

// MockStorage is a mock implementation of the Storage interface for testing.
type MockStorage struct {
	urlMappings map[string]string
	mu          sync.Mutex
}

// NewMockStorage creates a new MockStorage instance.
func NewMockStorage() *MockStorage {
	return &MockStorage{
		urlMappings: make(map[string]string),
	}
}

// StoreURLMapping stores a mapping between a short URL key and an original URL.
func (s *MockStorage) StoreURLMapping(shortURLKey, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urlMappings[shortURLKey] = originalURL
	return nil
}

// RetrieveOriginalURL retrieves the original URL associated with a short URL key.
func (s *MockStorage) RetrieveOriginalURL(shortURLKey string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	originalURL, exists := s.urlMappings[shortURLKey]
	if !exists {
		return "", ErrMappingNotFound
	}
	return originalURL, nil
}

// RetrieveShortURL retrieves the short URL associated with an original URL.
func (s *MockStorage) RetrieveShortURL(originalURL string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, url := range s.urlMappings {
		if url == originalURL {
			return key, nil
		}
	}
	return "", ErrMappingNotFound
}

// MockCache is a mock implementation of the Cache interface for testing.
type MockCache struct {
	cache map[string]string
	mu    sync.Mutex
}

// NewMockCache creates a new MockCache instance.
func NewMockCache() *MockCache {
	return &MockCache{
		cache: make(map[string]string),
	}
}

// Set stores a key-value pair in the mock cache.
func (c *MockCache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
	return nil
}

// Get retrieves a value from the mock cache by key.
func (c *MockCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.cache[key]
	// if !exists {
	// 	return "", ErrCacheMiss
	// }
	return value, exists
}
