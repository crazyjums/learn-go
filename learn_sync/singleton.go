package learn_sync

import "sync"

type Singleton struct {
	Name string
}

var singleton *Singleton
var syncOnce sync.Once

func GetSingleton() *Singleton {
	syncOnce.Do(func() {
		singleton = &Singleton{
			Name: "singleton",
		}
	})
	return singleton
}
