package main

import "sync/atomic"

func TestCas() {
	old := *new(*int32)
	atomic.CompareAndSwapInt32(old, 1, 2)
}
