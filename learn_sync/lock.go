package learn_sync

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	ThreadNum = 1000
)

var (
	rw sync.RWMutex
	co int
)

func WithNoLock() {
	var a = 0
	for i := 0; i < ThreadNum; i++ {
		go func(id int) {
			a += 1
			fmt.Println("i = ", id, ", a =", a)
		}(i)
	}
	fmt.Println("a=", a)
}

func WithMutexLock() {
	var a = 0
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		go func(id int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Println("i = ", id, ", a =", a)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("a=", a)
}

func WithWaitGroupMutex() {
	var a = 0
	var lock sync.Mutex
	var wg = sync.WaitGroup{}
	wg.Add(ThreadNum)
	for i := 0; i < ThreadNum; i++ {
		go func(id int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Println("i = ", id, ", a =", a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("a=", a)
}

func WithChanMutex() {
	var a = 0
	var lock sync.Mutex
	var c = make(chan bool, ThreadNum) //需要开辟内存空间，如果线程较多，不利于程序性能
	defer close(c)
	for i := 0; i < ThreadNum; i++ {
		go func(id int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Println("i = ", id, ", a =", a)
			c <- true //往通道写数据
		}(i)
	}

	for i := 0; i < ThreadNum; i++ {
		<-c //读取通道的数据
	}
	fmt.Println("a=", a)
}

func ReadWriteLock() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		read(i, ch)
	}
	for i := 0; i < 5; i++ {
		write(i, ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
}

func read(idx int, c chan struct{}) {
	rw.RLock()
	fmt.Println("goroutine 进入读操作... i = ", idx)
	v := co
	fmt.Println("goroutine 结束读操作，i = ", idx, ", v=", v)
	rw.RUnlock()
	c <- struct{}{}
}

func write(idx int, c chan struct{}) {
	rw.Lock()
	fmt.Println("goroutine ---进入写操作... i = ", idx)
	co = rand.Intn(1000)
	v := co
	fmt.Println("goroutine ---结束写操作，i = ", idx, ", new value =", v)
	rw.Unlock()
	c <- struct{}{}
}

func read2(i int) {
	fmt.Println(i, " is reading start")
	rw.RLock()
	fmt.Println(i, " is reading...")
	time.Sleep(time.Second)
	rw.RUnlock()
	fmt.Println(i, " is read over")
}

func write2(i int) {
	fmt.Println(i, " is write start")
	rw.Lock()
	fmt.Println(i, " is writing...")
	time.Sleep(time.Second)
	rw.Unlock()
	fmt.Println(i, " is write over")
}

func ReadAtSameTime() {
	go read2(1)
	go read2(2)
	time.Sleep(time.Second * 2)
}

func ReadWaitWrite() {
	go write2(1)
	go read2(2)
	go write2(3)
	time.Sleep(time.Second * 4)
}

func SyncMap() {
	var sm sync.Map

	sm.Store("name", "zhngsa")
	sm.Store("age", 100)

	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

type foo struct {
	name string
}

var f *foo
var once sync.Once

func SyncOnce(i int) *foo {
	once.Do(func() {
		f = &foo{}
		fmt.Println("addr=", f, ",i=", i)
	})

	return f
}

func main() {
	//得到汇编代码 go tool compile -N -l -S lock.go

	//WithNoLock()
	//WithMutexLock()
	//WithWaitGroupMutex()
	//WithChanMutex()
	//ReadWriteLock()
	//ReadAtSameTime()
	//ReadWaitWrite()
	//SyncMap()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			f := SyncOnce(i)
			log.Printf("f addr = %p\n", f)
			wg.Done()
		}(i + 1)
	}
	wg.Wait()
}
