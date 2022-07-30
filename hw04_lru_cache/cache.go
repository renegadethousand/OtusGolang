package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, ok := l.items[key]
	cahceItem := cacheItem{key, value}
	listItem := l.queue.PushFront(cahceItem)
	l.items[key] = listItem
	if l.queue.Len() > l.capacity {
		deleteItem := l.queue.Back()
		l.queue.Remove(deleteItem)
	}
	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	listItem, ok := l.items[key]
	if ok {
		l.queue.PushFront(listItem.Value)
		return listItem.Value.(cacheItem).value, ok
	}
	return listItem, ok
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}
