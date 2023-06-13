package main

import "fmt"

func cal() {
	var a int8 = -1
	var b int8 = -128 / a

	//int8 能表示的数字的范围是 [-2^7, 2^7-1]，即 [-128, 127]。-128 是无类型常量，
	//转换为 int8，再除以变量 -1，结果为 128，常量除以变量，结果是一个变量。
	//变量转换时允许溢出，符号位变为1，转为补码后恰好等于 -128。
	//
	//对于有符号整型，最高位是是符号位，计算机用补码表示负数。补码 = 原码取反加一。
	fmt.Println(b)//-128
}

func main() {
	const (
		a, b = 100, "abc"
		c, d
		e bool = false
	)
	fmt.Println(a, b, c, d, e)
	cal()

}
