package learn_sync

import (
	"log"
	"sync/atomic"
	"testing"
	"time"
)

func TestAddInt32(t *testing.T) {
	var a *int32
	atomic.AddInt32(a, 1)
}

func TestCompareAndSwap(t *testing.T) {
	spinLock()
}

func spinLock() {
	var num int64
	for atomic.CompareAndSwapInt64(&num, 10, 0) {
		log.Println("num = ", num)
	}
	time.Sleep(time.Nanosecond * 500)
}
