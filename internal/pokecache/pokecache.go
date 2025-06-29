package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry map[string]*cacheEntry
	mu    sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = &cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop() {

}

func NewCache(duration time.Duration) *Cache {
	return &Cache{
		entry: make(map[string]*cacheEntry),
	}
}
