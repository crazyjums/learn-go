package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int `json:"age" db:"a"`
}

func testStruct() {

	p := Person{"jums", 100}
	typeOfP := reflect.TypeOf(p)

	for i := 0; i < typeOfP.NumField(); i++ {
		fieldType := typeOfP.Field(i)
		fmt.Printf("%v, %v, %v \n", fieldType.Name, fieldType.Type, fieldType.Tag)

	}
}

func testPtr() {
	p := &Person{"jums", 100}
	typeOfP := reflect.TypeOf(p)

	fmt.Printf("name = %v, kind = %v \n", typeOfP.Name(), typeOfP.Kind())

	//取出指针地址中指向的元素信息
	typeOfPValue := typeOfP.Elem()
	fmt.Printf("elem name = %v, elem kind = %v", typeOfPValue.Name(), typeOfPValue.Kind())
}

func main() {
	var a int
	a = 1

	typeOfA := reflect.TypeOf(a)
	fmt.Printf("%v, %v \n", typeOfA.Name(), typeOfA.Kind())

	testStruct()
	testPtr()
}
