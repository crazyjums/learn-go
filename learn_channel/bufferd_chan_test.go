package main

import "testing"

func TestClosedChan(t *testing.T) {
	c := make(chan int)
	close(c)
	c <- 1
}
