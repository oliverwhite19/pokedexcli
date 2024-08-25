package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mx      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {

	c.mx.Lock()
	defer c.mx.Unlock()

	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reap(beforeTime time.Time) {

	c.mx.Lock()
	defer c.mx.Unlock()

	for key, value := range c.entries {
		if value.createdAt.Before(beforeTime) {
			delete(c.entries, key)
		}
	}

}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	now := time.Now()
	for range ticker.C {
		c.reap(now)
		now = time.Now()
	}

}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entries: make(map[string]cacheEntry),
		mx:      &sync.RWMutex{},
	}

	go newCache.reapLoop(interval)

	return newCache
}
