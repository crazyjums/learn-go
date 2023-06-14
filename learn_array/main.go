package main

import (
	"fmt"
	"math/big"
)

func Fibonacci(x int64) {
	var fib [1000]int64
	var i int64
	fib[0] = 0
	fib[1] = 1
	for i = 2; i < x; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println(fib)
	fmt.Println(fib[x-1])
}

func FibonacciBig(x int) {
	var fib [1000]*big.Int
	var i int
	fib[0] = big.NewInt(int64(1))
	fib[1] = big.NewInt(int64(1))
	for i = 2; i < x; i++ {
		t := fib[i].Add(fib[i-1], fib[i-2])
		fib[i] = t
	}
	fmt.Println(fib)
	fmt.Println(fib[x-1])
}

func main() {
	var arr = []int{6, 9, 3, 8, 7, 5}
	for i := range arr {
		fmt.Println(i)
	}
	//fmt.Println(arr)
	//
	//FibonacciBig(100)

	max := len(arr) - 1
	for i, e := range arr {
		if i == max {
			arr[0] += e
		} else {
			arr[i+1] += e
		}
	}
	fmt.Println(arr)

	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch int8(1) + 3 {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}

	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or26")
	}
}
