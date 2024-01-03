package main

import (
	"container/list"
	"fmt"
)

// LRUCache struct represents the least recently used chache.
type LRUCache struct {
	capacity int                      // maximum capacity of the cache
	cache    map[string]*list.Element // map to store chache entries for quick lookup.
	lruList  *list.List               // doubly linked list to maintain the order of recently used entries.
}

// entry struct represents a key-value piar in the cache.
type entry struct {
	key   string // key of the entry.
	value int    // value associated with the key.
}

// Get retrieves the value associatedwith the specified key from the cache
func (lru *LRUCache) Get(key string) int {
	if elem, ok := lru.cache[key]; ok {
		lru.lruList.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1 // key not found in the cache.
}

// Put inserts a key-value piar into the chache.
// if the cache is at its capacity, it evicts the Least recently used entry
func (lru *LRUCache) Put(key string, value int) {
	if elem, ok := lru.cache[key]; ok {
		// update existing entry
		elem.Value.(*entry).value = value
		lru.lruList.MoveToFront(elem)
	} else {
		// add new entry
		if len(lru.cache) >= lru.capacity {
			// Evict the Least recently used entry
			back := lru.lruList.Back()
			if back != nil {
				delete(lru.cache, back.Value.(*entry).key)
				lru.lruList.Remove(back) // Remove the least recently used entry from the list.
			}
		}
		newEntry := &entry{key, value}
		newElem := lru.lruList.PushFront(newEntry)
		lru.cache[key] = newElem // update the map with the new entry
	}
}

func newLRU(size int) *LRUCache {
	return &LRUCache{
		capacity: size,
		cache:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
}

func main() {
	x := newLRU(2)
	x.Put("1", 1)           // 1
	x.Put("2", 2)           // 1, 2
	fmt.Println(x.Get("1")) // ok
	x.Put("3", 3)           // 1, 3
	fmt.Println(x.Get("2")) // not ok
	fmt.Println(x.Get("3")) // ok
}
