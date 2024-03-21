package lru

import (
	"errors"
	"time"
)

var ErrKeyNotFound = errors.New("key not found")

type lruItem struct {
	key          string
	value        any
	lastTimeUsed time.Time
}

type LRUCache struct {
	capacity int
	items    map[string]*lruItem
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*lruItem),
	}
}

func (lru *LRUCache) Get(key string) (any, error) {
	item, ok := lru.items[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	item.lastTimeUsed = time.Now()
	return item.value, nil
}

func (lru *LRUCache) Set(key string, value any) {
	if len(lru.items) == lru.capacity {
		oldestKey := ""
		oldestTime := time.Now()
		for i := range lru.items {
			if lru.items[i].lastTimeUsed.Before(oldestTime) {
				oldestTime = lru.items[i].lastTimeUsed
				oldestKey = i
			}
		}
		delete(lru.items, oldestKey)
	}

	lru.items[key] = &lruItem{key: key, value: value, lastTimeUsed: time.Now()}
}

func (lru *LRUCache) Len() int {
    return len(lru.items)
}
