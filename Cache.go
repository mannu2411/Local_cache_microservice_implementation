package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type CachedData struct {
	Data              string `json:"email"`
	expireAtTimestamp time.Time
}

type Caches struct {
	mu    sync.Mutex
	cache map[string]CachedData
}

func NewCacheProvider(interval time.Duration) CacheProvider {
	c := Caches{
		mu:    sync.Mutex{},
		cache: make(map[string]CachedData),
	}
	go func(cleanupInterval time.Duration) {
		c.cleanupLoop(cleanupInterval)
	}(interval)
	return &c
}

//func (cache *Caches) NewLocalCache(cleanupInterval time.Duration) *Caches {
//	cache = &Caches{
//		mu:    sync.Mutex{},
//		cache: make(map[string]CachedData),
//	}
//	go func(cleanupInterval time.Duration) {
//		cache.cleanupLoop(cleanupInterval)
//	}(cleanupInterval)
//	return cache
//}

func (cache *Caches) AddToCache(key, data []byte) {
	c := CachedData{Data: string(data), expireAtTimestamp: time.Now().Add(time.Minute * 15)}
	// Local Cache storing data for 15 minutes
	cache.mu.Lock()
	cache.cache[string(key)] = c
	cache.mu.Unlock()
}
func (cache *Caches) GetCache(key []byte) (string, error) {
	c, ok := cache.cache[string(key)]
	if !ok {
		return "", fmt.Errorf("unable to get data from cache for key: %v", key)
	}
	if c.expireAtTimestamp.Before(time.Now()) {
		delete(cache.cache, string(key))
		return "", fmt.Errorf("unable to get data from cache for key: %v", key)
	}
	return c.Data, nil
}

func (cache *Caches) GetMap() map[string]CachedData {
	return cache.cache
}
func (cache *Caches) delete(key []byte) {
	delete(cache.cache, string(key))
	log.Printf("%v was deleted from cache", key)
}

func (cache *Caches) cleanupLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		for uid, cu := range cache.cache {
			if cu.expireAtTimestamp.Before(time.Now()) {
				delete(cache.cache, uid)
				log.Printf("%v was deleted from cache", uid)
			}
		}
	}
}
