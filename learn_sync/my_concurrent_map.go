package learn_sync

import (
	"sync"
)

type ConcurrentMap struct {
	mu    sync.RWMutex
	items map[string]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		items: make(map[string]interface{}),
	}
}

func (cm *ConcurrentMap) Set(key string, value interface{}) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.items[key] = value
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	value, ok := cm.items[key]
	return value, ok
}

func (cm *ConcurrentMap) Delete(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.items, key)
}

func (cm *ConcurrentMap) Len() int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return len(cm.items)
}
