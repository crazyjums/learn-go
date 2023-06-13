package main

import "fmt"

type Person interface {
	getName() string
	getAge() int
}

type Student struct {
	name string
	age  int
}

// InitAnonymousStruct 初始化匿名结构体
func InitAnonymousStruct() {
	a := struct {
		name string
		age  int
		tall float32
	}{
		"lisi",
		14,
		173.2,
	}

	b := &struct {
		name string
		age  int
		tall float32
	}{
		"lisi",
		14,
		173.2,
	}

	fmt.Println("初始化匿名结构体")
	fmt.Println(a, b)
}

// InitStructByLiteral 通过字面量初始化
func InitStructByLiteral() {
	stu := Student{
		name: "zhangsan",
		age:  14,
	}

	stu3 := Student{
		name: "hahh",
	}

	fmt.Println("通过字面量初始化")
	fmt.Println(stu, stu3)
}

// InitStructByVar 通过var初始化
func InitStructByVar() {
	var stu Student
	stu.name = "wangwu"
	stu.age = 16

	fmt.Println("通过var初始化")
	fmt.Println(stu)
}

// InitStructByNew 通过new初始化
func InitStructByNew() {
	stu := new(Student)
	stu.name = "wangwu"
	stu.age = 18

	//这是new的一种简写，底层还是用的new关键字初始化的结构体
	stu2 := &Student{
		"zhangsan",
		15,
	}

	fmt.Println("通过new初始化")
	fmt.Println(stu, stu2)
}

func main() {
	InitAnonymousStruct()
	InitStructByLiteral()
	InitStructByVar()
	InitStructByNew()
	fmt.Println(1&100==0)
}
