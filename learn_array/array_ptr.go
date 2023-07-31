package learn_array

import "fmt"

func ArrayPtr() {
	// 声明一个数组
	arr := [5]int{1, 2, 3, 4, 5}

	// 获取数组的指针
	ptr := &arr

	// 使用指针访问数组元素
	fmt.Println((*ptr)[0]) // 输出: 1
	fmt.Println((*ptr)[1]) // 输出: 2
	fmt.Println((*ptr)[2]) // 输出: 3
}
