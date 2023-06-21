package learn_sync

import "sync"

func Map() {
	sm := sync.Map{}
	sm.Store("key", "value")
}
