package main

import (
	"time"
)

type CacheProvider interface {
	AddToCache(key, data []byte)
	GetCache(key []byte) (string, error)
	delete(key []byte)
	cleanupLoop(interval time.Duration)
	GetMap() map[string]CachedData
	//NewLocalCache(cleanupInterval time.Duration) *Caches
}
