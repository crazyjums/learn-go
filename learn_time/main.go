package main

import (
	"fmt"
	"time"
)

func main() {
	curr := time.Now().Format("010215")
	if curr >= "060114" && curr <= "060522" {
		fmt.Printf("unsafe  curr=%v", curr)
	} else {
		fmt.Printf("safe curr=%v", curr)
	}
	//hb := time.NewTicker(5 * time.Second)
	//fmt.Printf("time=%v", time.Now().Unix())
	//defer hb.Stop()
	//select {
	//case <-hb.C:
	//	fmt.Printf("time=%v", time.Now().Unix())
	//}
}
