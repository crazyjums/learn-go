package learn_sync

import (
	"sync"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	cm := NewConcurrentMap()

	// Set values concurrently
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			cm.Set(string(rune(key)), key)
		}(i)
	}
	wg.Wait()

	// Get values concurrently
	wg = sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			value, ok := cm.Get(string(rune(key)))
			if ok {
				println(value.(int))
			}
		}(i)
	}
	wg.Wait()

	// Delete values concurrently
	wg = sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			cm.Delete(string(rune(key)))
		}(i)
	}
	wg.Wait()

	println(cm.Len()) // Output: 0
}
