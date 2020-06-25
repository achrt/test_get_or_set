package main

import (
	"go_test/cache"
	"math/rand"
	"strconv"
)

func main() {
	memCache := cache.NewInMemoryCache()
	f := func() string {
		rand := rand.Intn(10)
		return strconv.Itoa(rand)
	}

	for i := 0; i < 1000; i++ {
		go memCache.GetOrSet(strconv.Itoa(rand.Intn(10)), f)
	}
}
