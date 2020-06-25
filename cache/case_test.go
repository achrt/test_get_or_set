package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)


func TestGetOrSet(t *testing.T) {
	var count int64 = 0
	memCache := NewInMemoryCache()
	f := func() string {
		atomic.AddInt64(&count, 1)
		return strconv.Itoa(randomInt(1, 10))
	}

	for i := 0; i < 100000; i++ {
		go memCache.GetOrSet(strconv.Itoa(randomInt(1, 10)), f)
	}
	countFinal := atomic.LoadInt64(&count)
	if countFinal > 10 {
		t.Error(fmt.Sprintf("Counter limit! Value: %d", countFinal))
	}
}

func TestGet(t *testing.T) {
	memCache := NewInMemoryCache()

	for i := 0; i < 1000; i++ {
		go memCache.Get(strconv.Itoa(randomInt(1, 10)))
	}
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
