package main

import "testing"

func TestFetchByOne(t *testing.T) {
	FetchWith1Goroutine()
}

func TestFetchByTen(t *testing.T) {
	FetchWith10Goroutine()
}

func TestWgWithDeadLock(t *testing.T) {
	wgWithDeadLock()
}

func TestWgWithoutDeadLock(t *testing.T) {
	wgWithoutDeadLock()
}

func TestWgPage(t *testing.T) {
	wgPage()
}

func TestWgPage2(t *testing.T) {
	wgPage2()
}
