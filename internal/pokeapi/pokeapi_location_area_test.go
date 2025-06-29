package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockCache implements the CacheInterface for testing
type MockCache struct {
	data map[string][]byte
}

func NewMockCache() *MockCache {
	return &MockCache{
		data: make(map[string][]byte),
	}
}

func (m *MockCache) Add(key string, val []byte) {
	m.data[key] = val
}

func (m *MockCache) Get(key string) ([]byte, bool) {
	val, ok := m.data[key]
	return val, ok
}

func TestGetLocations(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a mock response
		response := LocationResponse{
			Count: 2,
			Next:  "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			Prev:  "",
			Results: []LocationArea{
				{Name: "canalave-city-area", Url: "https://pokeapi.co/api/v2/location-area/1/"},
				{Name: "eterna-city-area", Url: "https://pokeapi.co/api/v2/location-area/2/"},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Test with no cache
	t.Run("No cache", func(t *testing.T) {
		config := &Config{}
		locations, err := GetLocations(config, server.URL)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(locations) != 2 {
			t.Errorf("Expected 2 locations, got %d", len(locations))
		}

		if locations[0].Name != "canalave-city-area" {
			t.Errorf("Expected first location to be 'canalave-city-area', got '%s'", locations[0].Name)
		}

		if config.Next != "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20" {
			t.Errorf("Expected Next to be set correctly, got '%s'", config.Next)
		}

		if config.Prev != "" {
			t.Errorf("Expected Prev to be empty, got '%s'", config.Prev)
		}
	})

	// Test with cache
	t.Run("With cache", func(t *testing.T) {
		cache := NewMockCache()
		config := &Config{
			Cache: cache,
		}

		// First call should hit the API
		locations, err := GetLocations(config, server.URL)
		if err != nil {
			t.Fatalf("Expected no error on first call, got %v", err)
		}

		if len(locations) != 2 {
			t.Errorf("Expected 2 locations on first call, got %d", len(locations))
		}

		// Second call should use the cache
		locations, err = GetLocations(config, server.URL)
		if err != nil {
			t.Fatalf("Expected no error on second call, got %v", err)
		}

		if len(locations) != 2 {
			t.Errorf("Expected 2 locations on second call, got %d", len(locations))
		}

		// Verify cache was used by checking if the data is in the mock cache
		_, ok := cache.Get(server.URL)
		if !ok {
			t.Errorf("Expected data to be in cache")
		}
	})

	// Test error handling for invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		config := &Config{}
		_, err := GetLocations(config, "http://invalid-url")

		if err == nil {
			t.Errorf("Expected error for invalid URL, got nil")
		}
	})
}
