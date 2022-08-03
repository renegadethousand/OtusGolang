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

	listItem, ok := l.items[key]
	if ok {
		listItem.Value = cacheItem{key, value}
		l.queue.MoveToFront(listItem)
		return true
	}
	l.items[key] = l.queue.PushFront(cacheItem{key: key, value: value})
	if l.queue.Len() > l.capacity {
		deleteItem := l.queue.Back()
		l.queue.Remove(deleteItem)
		delete(l.items, deleteItem.Value.(cacheItem).key)
	}
	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	listItem, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(listItem)
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
