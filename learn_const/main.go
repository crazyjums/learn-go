package main

import "fmt"

func cal()  {
	const a int8 = -1
	var b int8 = -128 / a

	//-128 和 a 都是常量，在编译时求值，-128 / a = 128，两个常量相除，
	//结果也是一个常量，常量类型转换时不允许溢出，因而编译失败。
	fmt.Println(b)
}

func main() {
	const a = 100
	var b int32 = a
	fmt.Println(a, b)

	const c = 100
	var d int64 = c
	fmt.Println(c, d)

	const e = 100
	var f float32 = e
	fmt.Println(e, f)

	const g = 100
	var h float32 = g
	fmt.Println(g, h)

	const i int32 = 100
	var j = i
	fmt.Println(i, j)

	//const x int32 = 100
	//var y int64 = x //cannot use x (type int32) as type int64 in assignment
	//fmt.Println(x, y)

	cal()
}
