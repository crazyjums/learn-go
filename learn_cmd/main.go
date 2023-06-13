package main

import (
	"flag"
	"fmt"
)

var name string

func init()  {
	//接收命令行动态传入参数，第一个参数指定变量接收，第二个参数为该参数的名字，第三个参数为默认值，第四个为用法提示
	flag.StringVar(&name, "name", "default name", "get a name")

	name1 := flag.String("name1", "jums", "get a greeting object")
	fmt.Println(*name1)
}

func main() {
	flag.Parse() //用于真正解析命令参数，并把它们的值赋给相应的变量。
	fmt.Printf("hello, %s \n", name)
	//go build main.go
	//main.exe -help
	//main.exe -name="new name"
}
