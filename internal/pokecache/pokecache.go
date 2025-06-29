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
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	timer := time.NewTimer(interval)
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				c.mu.Lock()
				now := time.Now()
				threshold := now.Add(-interval)
				for key, entry := range c.entry {
					if entry.createdAt.Before(threshold) {
						delete(c.entry, key)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
}

func NewCache(duration time.Duration) *Cache {
	cache := &Cache{
		entry: make(map[string]*cacheEntry),
	}
	cache.reapLoop(duration)
	return cache
}
