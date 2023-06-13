package main

import "fmt"

func processError()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("r: %v \n", r)
		}
	}()

	b := 0
	a := 1 / b
	fmt.Println(a)
}

func main() {
	processError()
	fmt.Println("gg")
}
