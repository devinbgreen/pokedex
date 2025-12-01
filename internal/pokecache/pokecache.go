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
	critical sync.Mutex
}

func NewCache(interval time.Duration) {
	//
}

func (c *Cache) add(key string, val []byte) error {
	//
	return nil
}

func (c *Cache) get(key string) ([]byte, bool) {
	//
	return nil, true
}

func (c *Cache) reapLoop() {

}
