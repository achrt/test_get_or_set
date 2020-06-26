package cache

import (
	"sync"
)

type (
	Key   = string
	Value = string
)

type Cache interface {
	GetOrSet(key Key, valueFn func() Value) Value
	Get(key Key) (Value, bool)
	Set(key Key, value Value)
}

type InMemoryCache struct {
	dataMutex sync.RWMutex
	data      map[Key]Value
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[Key]Value),
	}
}

func (cache *InMemoryCache) Get(key Key) (Value, bool) {
	cache.dataMutex.RLock()
	defer cache.dataMutex.RUnlock()
	return cache.get(key)
}

func (cache *InMemoryCache) Set(key Key, value Value) {
	cache.dataMutex.Lock()
	defer cache.dataMutex.Unlock()
	cache.set(key, value)
}

func (cache *InMemoryCache) GetOrSet(key Key, valueFn func() Value) Value {
	cache.dataMutex.Lock()
	defer cache.dataMutex.Unlock()
	value, found := cache.get(key)
	if found {
		return value
	}
	findValue := valueFn()
	cache.set(key, findValue)
	return findValue
}

func (cache *InMemoryCache) set(key Key, value Value) {
	cache.data[key] = value
}
func (cache *InMemoryCache) get(key Key) (Value, bool) {
	value, found := cache.data[key]
	return value, found
}
