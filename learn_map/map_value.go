package learn_map

import "fmt"

func MapValueFor() {
	var m = make(map[string]*Student)
	stus := []Student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	fmt.Printf("%p %p %p \n", &stus[0], &stus[1], &stus[2])

	for _, stu := range stus {
		//这里的 stu 是一个临时变量，每次循环都会被重新赋值，所以最终 map 中的所有元素的指针都指向了 stu 的地址，
		//而 stu 的地址是不变的，所以最终 map 中的所有元素的值都是最后一次循环的结果。
		m[stu.Name] = &stu
		fmt.Printf("%p, %v\n", &stu, m)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.Name) //输出结果为：zhou => wang, li => wang, wang => wang
	}
}

func MapValueForV2() {
	var m = make(map[string]*Student)
	stus := []Student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	fmt.Printf("%p %p %p \n", &stus[0], &stus[1], &stus[2])

	for _, stu := range stus {
		t := stu
		m[stu.Name] = &t
		fmt.Printf("%p, %v\n", &stu, m)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.Name) // 输出结果为：zhou => zhou, li => li, wang => wang
	}
}
