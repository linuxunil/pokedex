package pokedex

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	lock  *sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	ticker := time.NewTicker(interval)
	cache := Cache{make(map[string]cacheEntry), &sync.Mutex{}}
	go cache.reapLoop(ticker, interval)
	return cache

}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	entry := cacheEntry{time.Now(), val}
	c.entry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	entry, ok := c.entry[key]
	if ok {
	} else {
	}
	return entry.val, ok
}

func (c *Cache) reapLoop(ticker *time.Ticker, interval time.Duration) {
	for range ticker.C {
		for i := range c.entry {
			c.lock.Lock()
			age := time.Now().Sub(c.entry[i].createdAt).Seconds()
			if age > interval.Seconds() {
				delete(c.entry, i)
			}
			c.lock.Unlock()
		}
	}
}
