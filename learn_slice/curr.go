package learn_slice

import (
	"fmt"
	"sync"
)

type SafeSlice struct {
	slice []int
	mu    sync.RWMutex
}

func (ss *SafeSlice) Append(value int) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.slice = append(ss.slice, value)
}

func (ss *SafeSlice) Get(index int) int {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	return ss.slice[index]
}

func (ss *SafeSlice) Len() int {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	return len(ss.slice)
}

func (ss *SafeSlice) Print() {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	fmt.Println(ss.slice)
}
