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
	stuff    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		stuff:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	/*
			_, exists := c.stuff[key]
		    if exists {
		        // maybe overwrite, maybe early-return – your choice for this cache
		    }*/
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.stuff[key] = entry
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.stuff[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.stuff {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.stuff, key)
			}
		}
		c.mu.Unlock()
	}
}
