package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mux      sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type keyValue struct {
	Key   Key
	Value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mux.Lock()
	defer l.mux.Unlock()
	item, ok := l.items[key]
	newValue := &keyValue{
		Key:   key,
		Value: value,
	}
	if ok {
		item.Value = newValue
		l.queue.MoveToFront(item)
		return true
	}
	if l.capacity == l.queue.Len() {
		backItem := l.queue.Back().Value.(*keyValue)
		delete(l.items, backItem.Key)
		l.queue.Remove(l.queue.Back())
	}
	item = l.queue.PushFront(newValue)
	l.items[key] = item
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mux.Lock()
	defer l.mux.Unlock()
	item, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(item)
		itemValue := item.Value.(*keyValue)
		return itemValue.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
