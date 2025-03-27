package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mux     *sync.Mutex
	Entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{mux: &sync.Mutex{}, Entries: make(map[string]cacheEntry)}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.Entries[key] = cacheEntry{time.Now(), val}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		for key, entry := range c.Entries {
			if entry.createdAt.Add(interval).Before(time.Now()) {
				c.mux.Lock()
				delete(c.Entries, key)
				c.mux.Unlock()
			}
		}
	}
}
