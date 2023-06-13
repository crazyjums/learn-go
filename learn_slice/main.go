package main

import "fmt"

func main() {
	a := make([]int, 5) // 指定长度为5，容量没有设置，则和长度相同：len=5 cap=5
	printSlice("a", a)
	b := make([]int, 0, 5) // 指定长度为0，容量为5：len=0 cap=5
	printSlice("b", b)
	c := b[:2] // 切片时长度为下表差，容量计算时需要考虑左下标开始位置，这里的左下标从0开始：len=2 cap=5
	printSlice("c", c)
	d := c[2:5] // 切片时长度为下表差，容量计算时需要考虑左下标开始位置，这里的左下标从2开始：len=3 cap=5
	printSlice("d", d)

	a1 := []int{1, 2, 3, 4, 5}
	printSlice("a1", a1)
	a2 := modArray(a1)
	printSlice("a2", a2)
	printSlice("a1", a1)
}

func modArray(a []int) []int {
	a[0] = 11
	return a
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
