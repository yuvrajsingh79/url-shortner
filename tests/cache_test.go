package tests

import (
	"interview/url-shortner/controller"
	"testing"
)

func TestRedisCache_SetAndGet(t *testing.T) {
	// Initialize a Redis cache for testing.
	redisCache := controller.NewMockCache()

	t.Run("SetAndGet_ValidData_ReturnsCorrectValue", func(t *testing.T) {
		// Test setting a key-value pair in the cache and retrieving it.
		key := "testkey"
		value := "testvalue"

		err := redisCache.Set(key, value)
		if err != nil {
			t.Fatalf("Error setting value: %v", err)
		}

		retrievedValue, found := redisCache.Get(key)
		if !found {
			t.Fatalf("Key %s not found in cache", key)
		}

		if retrievedValue != value {
			t.Errorf("Expected value %s; got %s", value, retrievedValue)
		}
	})

	// Add more test cases for different scenarios.

	// Cleanup or reset the mock cache if needed.
}
