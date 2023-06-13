package main

import (
	"fmt"
	"time"
)

func loop(i int) {
	t := 1
	for {
		t++
		fmt.Printf("t:= %v, (%v) \n", t, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//go loop(1)
	//go loop(2)
	//
	//var i string
	//_, err := fmt.Scanln(&i)
	//if err != nil {
	//	return
	//}

	//c := make(chan int, 1)
	//for {
	//	select {
	//	case a:=<-c:
	//		fmt.Println("<-c",a)
	//	case c <- 1:
	//		fmt.Println("c<-1")
	//	}
	//}
	TestRwMu()
}
