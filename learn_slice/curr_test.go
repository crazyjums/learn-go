package learn_slice

import (
	"sync"
	"testing"
)

func TestSafeSlice_Append(t *testing.T) {
	s := &SafeSlice{}
	count := 10
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < count; i++ {
			s.Append(i)
		}
		wg.Done()
	}()

	go func() {
		for i := count; i < count*2; i++ {
			s.Append(i)
		}
		wg.Done()
	}()

	wg.Wait()
	t.Log(s.Len())
	s.Print()
}
