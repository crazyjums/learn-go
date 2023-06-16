package learn_sync

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	flag *int32
}

func NewSpinLock() *SpinLock {
	return &SpinLock{flag: new(int32)}
}

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(s.flag, 0, 1) {
		//让出cpu，并不是暂停当前goroutine的执行，而是去执行其他等待的任务
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32(s.flag, 0)
}
