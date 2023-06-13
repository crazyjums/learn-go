package main

import (
	"fmt"
	"sync"
	"time"
)

type CounterRw struct {
	mu    sync.RWMutex
	count int
}

func (c *CounterRw) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *CounterRw) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestRwMu() {
	var c CounterRw
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(c.Count())
			time.Sleep(time.Millisecond)
		}()
	}

	for {
		c.Incr()
		time.Sleep(time.Second)
	}
}
