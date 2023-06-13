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
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)

	FibonacciBig(100)
}
