package cache

import (
	"sync"
)

/*
Дано:
	InMemoryCache - потоко-безопасная реализация Key-Value кэша, хранящая данные в оперативной памяти
Задача:
	1. Реализовать метод GetOrSet, предоставив следующие гарантии:
		- Значение каждого ключа будет вычислено ровно 1 раз
		- Конкурентные обращения к существующим ключам не блокируют друг друга
	2. Покрыть его тестами, проверить метод 1000+ горутинами

Что должно быть в тесте:
1000 горутин
Каждая обращается к случайному ключу от 1 до 10
Когда все горутины отработали, valueFn должна быть вызвана ровно 10 раз
*/

// ----------------------------------------------

type (
	Key   = string
	Value = string
)

type Cache interface {
	GetOrSet(key Key, valueFn func() Value) Value
	Get(key Key) (Value, bool)
	Set(key Key, value Value)
}

// ----------------------------------------------

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

// GetOrSet возвращает значение ключа в случае его существования.
// Иначе, вычисляет значение ключа при помощи valueFn, сохраняет его в кэш и возвращает это значение.

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
