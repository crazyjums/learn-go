package main

import "sync"

func TestMutex() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
}
