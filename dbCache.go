package main

import (
	"fmt"

	lru "github.com/flyaways/golang-lru"
)

// functions below implement Least Recently Used (LRU) caching strategy
// It is used in cases when we assume that the most liekly user to return
// is the one that was fetched the most recently
// (I wasn't sure which would make more sense in the use case of activity
// logging beetwen this and LFU so I choose what seemed easier to do)

var (
	cache               lru.LRUCache
	err                 error
	name                string
	ok                  bool
	sqlRequestCount     int
	cacheRetrivalsCount int
)

func postgresql(id int) (int, string) {
	//Lets pretend that entire human population is named Adam :)
	sqlRequestCount += 1
	return id, "Adam"
}

func fetchUserById(id int) (int, string) {
	//first, attempt to retreive the user from cache
	name, ok := cache.Get(id)
	fmt.Println(name, ok)

	if ok {
		cacheRetrivalsCount += 1

		name, ok := name.(string)
		if ok {
			return id, name
		}
	}
	return postgresql(id)
}

func main() {
	cache, err = lru.New(100)
	// worst case scenario for lru
	for i := 1; i <= 100; i++ {
		for id := 1; id <= 100; id++ {
			fetchUserById(id)
		}
	}

	fmt.Println(fmt.Sprintf("records retreived from sql: %d | records retreived from cache: %d", sqlRequestCount, cacheRetrivalsCount))
}
