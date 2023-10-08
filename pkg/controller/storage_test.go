package controller

import (
	"testing"
)

func TestMemoryStorage_StoreAndRetrieve(t *testing.T) {
	// Initialize a memory storage for testing.
	memoryStorage := NewMemoryStorage()

	t.Run("StoreAndRetrieve_ValidData_ReturnsCorrectValue", func(t *testing.T) {
		// Test storing and retrieving a URL mapping.
		shortURLKey := "abc123"
		originalURL := "https://example.com"

		err := memoryStorage.StoreURLMapping(shortURLKey, originalURL)
		if err != nil {
			t.Fatalf("Error storing URL mapping: %v", err)
		}

		retrievedURL, err := memoryStorage.RetrieveOriginalURL(shortURLKey)
		if err != nil {
			t.Fatalf("Error retrieving URL mapping: %v", err)
		}

		if retrievedURL != originalURL {
			t.Errorf("Expected URL %s; got %s", originalURL, retrievedURL)
		}
	})

	// Add more test cases for different scenarios.

	// Cleanup or reset the memory storage if needed.
}
