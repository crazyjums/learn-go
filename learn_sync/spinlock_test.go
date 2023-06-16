package learn_sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestSpinLock(t *testing.T) {
	spinLock := NewSpinLock()
	wg := sync.WaitGroup{}

	var i, num int
	for i = 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			spinLock.Lock()
			defer func() {
				spinLock.Unlock()
				wg.Done()
			}()
			num++
		}()
	}
	wg.Wait()
	fmt.Println("num = ", num)
}
