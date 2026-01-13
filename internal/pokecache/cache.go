package pokecache

import (
	"sync"
	"time"
)


type Cache struct {
	mux *sync.Mutex
	pokeCache map[string]CacheEntry
}


type CacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache(interval time.Duration) Cache {
	c := Cache{
		mux:       &sync.Mutex{},
		pokeCache: make(map[string]CacheEntry),
	}
	
	go c.reapCache(interval)
	return c
}


func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.pokeCache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.pokeCache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}


func (c *Cache) reapCache(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		defer c.mux.Unlock()
		for key, entry := range c.pokeCache {
			if time.Since(entry.createdAt) > interval {
				delete(c.pokeCache, key)
			}
		}
	}
}