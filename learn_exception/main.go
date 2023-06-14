package main

import "fmt"

func processError() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("r: %v \n", r)
		}
	}()

	b := 0
	a := 1 / b
	fmt.Println(a)
}

func TestDefer() {
	panic("test defer")
	//panic发生时，之后的代码不会执行，所有这个defer中的代码不会被执行
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover")
		}
		fmt.Println("defer")
	}()
}

func TestDefer2() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover")
		}
		fmt.Println("defer")
	}()
	panic("test defer")
}

func TestMultiDefer() {
	//如果一个函数中有多个defer，他们的执行顺序和他们的出现顺序是相反的
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		fmt.Println("defer2")
	}()

	panic("ttt")
}

func TestDeferInFor() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}

func main() {
	processError()
	fmt.Println("gg")
	//TestDefer2()
	TestDeferInFor()
}
