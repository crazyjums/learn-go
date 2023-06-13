package main

import "time"

func TestSelect() {
	c := make(chan int)
	defer close(c)

	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
	}()

	select {
	case v := <-c:
		println(v)
	case <-time.After(3 * time.Second):
		println("timeout")
	}
}
