package main

import "testing"

func TestT(t *testing.T) {
	TestDefer()
	TestDefer2()
	TestMultiDefer()
	TestDeferInFor()
}
