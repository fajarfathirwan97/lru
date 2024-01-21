package list

import (
	"sync"
)

type LRU struct {
	cache         *sync.Map
	invertedCache *sync.Map
	maxSize       int
	len           int64
	list          *DoublyLinkedList
}

func NewLRU(maxSize int) *LRU {
	return &LRU{
		maxSize:       maxSize,
		cache:         &sync.Map{},
		invertedCache: &sync.Map{},
		list:          &DoublyLinkedList{}}
}
func (l *LRU) Len() int64 {
	return l.list.len
}

func (l *LRU) Add(key, data string) string {
	v, ok := l.cache.Load(key)
	if ok {
		node := v.(*Node)
		node.Data = data
		l.cache.Store(key, data)
		l.invertedCache.Store(key, data)
		return node.Data
	}
	l.len++
	node := l.list.Prepend(data)
	if l.len > int64(l.maxSize) {
		l.list.RemoveTail()
		l.cache.Delete(key)
		l.invertedCache.Delete(key)
		l.len--
	}
	l.cache.Store(key, node)
	l.invertedCache.Store(node, key)
	return node.Data
}

func (l LRU) Get(key string) string {
	v, ok := l.cache.Load(key)
	if !ok {
		return ""
	}
	node := v.(*Node)
	if node.prev == nil {
		return node.Data
	} else {
		l.list.RemoveAtNode(node)
		l.list.Prepend(node.Data)
		return node.Data
	}
}
