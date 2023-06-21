package learn_sync

import "sync"

type Person struct {
	Name string
}

func SyncPool() {
	pool := sync.Pool{
		New: func() interface{} {
			return &Person{}
		},
	}

	p := pool.Get().(*Person)
	p.Name = "张三"
	pool.Put(p)
}
