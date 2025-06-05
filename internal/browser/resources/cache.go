package resources

import (
	"sync"
)

// MemoryCache is a simple in-memory implementation of ResourceCache
type MemoryCache struct {
	mu    sync.RWMutex
	cache map[string][]byte
}

// NewMemoryCache creates a new memory-based resource cache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: make(map[string][]byte),
	}
}

// Get retrieves cached content for the given URL
func (mc *MemoryCache) Get(url string) ([]byte, bool) {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	content, found := mc.cache[url]
	if found {
		// Return a copy to prevent modification
		result := make([]byte, len(content))
		copy(result, content)
		return result, true
	}
	return nil, false
}

// Set stores content for the given URL
func (mc *MemoryCache) Set(url string, data []byte) {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	// Store a copy to prevent external modification
	content := make([]byte, len(data))
	copy(content, data)
	mc.cache[url] = content
}

// Clear removes all cached content
func (mc *MemoryCache) Clear() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.cache = make(map[string][]byte)
}

// Size returns the number of cached items
func (mc *MemoryCache) Size() int {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	return len(mc.cache)
}

// Has checks if a URL is cached without retrieving the content
func (mc *MemoryCache) Has(url string) bool {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	_, found := mc.cache[url]
	return found
}

// Delete removes a specific URL from the cache
func (mc *MemoryCache) Delete(url string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	delete(mc.cache, url)
}

// Keys returns all cached URLs
func (mc *MemoryCache) Keys() []string {
	mc.mu.RLock()
	defer mc.mu.RUnlock()

	keys := make([]string, 0, len(mc.cache))
	for url := range mc.cache {
		keys = append(keys, url)
	}
	return keys
}
