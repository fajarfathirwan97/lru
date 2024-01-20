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
	return l.len
}

func (l *LRU) Add(key, data string) {
	v, ok := l.cache.Load(key)
	if ok {
		node := v.(*Node)
		node.Data = data
		l.cache.Store(key, data)
		l.invertedCache.Store(key, data)
		return
	}
	node := l.list.Append(data)
	l.len++
	if l.len > int64(l.maxSize) {
		l.list.Remove(l.list.Tail.Data)
		l.cache.Delete(key)
		l.invertedCache.Delete(key)
		l.len--
	}
	l.cache.Store(key, node)
	l.invertedCache.Store(node, key)
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
		l.list.Append(node.Data)
		return node.Data
	}
}
