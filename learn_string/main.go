package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"unicode/utf8"
)

func main() {

	//---字符串比较---
	fmt.Println("ca" > "b") //false
	fmt.Println("b" > "a")  //true

	//---字符串拼接---
	//因为编译器会在行尾自动补全分号，所以拼接字符串用的加号“+”必须放在第一行末尾。
	str := "Beginning of the string " +
		"second part of the string"
	fmt.Println(str) //Beginning of the string second part of the string

	//---计算字符串的长度---
	//ASCII 字符串长度使用 len() 函数。
	//Unicode 字符串长度使用 utf8.RuneCountInString() 函数。
	str1 := "I am a man;"
	str2 := "很好aa"
	//len() 函数的返回值的类型为 int，表示字符串的 ASCII 字符个数或字节长度
	//如果字符串中有中文，返回的结果可能会不是我们想要的结果
	//因为在go中，字符串都是用utf8编码，一个中文字占3个字符
	fmt.Println(len(str1), len(str2)) //11 6
	//utf8.RuneCountInString会将字符串中的中文当成一个"字符"来统计
	fmt.Println(utf8.RuneCountInString(str1), utf8.RuneCountInString(str2)) //11 2
	for i, w := 0, 0; i < len(str2); i += w {
		runeV, width := utf8.DecodeRuneInString(str2[i:])
		fmt.Printf("%#U %d \n", runeV, width)
		w = width
	}

	//---字符串遍历---
	//ASCII 字符串遍历直接使用下标。
	//Unicode 字符串遍历用 for range。
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c %d \n", str2[i], str2[i])
	}

	for _, v := range str2 {
		fmt.Printf("%c %d \n", v, v)
	}

	//---字符串截取---
	//strings.Index：正向搜索子字符串。
	//strings.LastIndex：反向搜索子字符串。
	fmt.Println(strings.LastIndex(str2, "a")) //7
	fmt.Println(strings.Index(str2, "a"))     //6

	var byteStr []byte
	byteStr = append(byteStr, 'a')
	byteStr = append(byteStr, 'b')
	fmt.Println(string(byteStr))
	byteStr[0] = 'c'
	fmt.Println(string(byteStr))

	str3 := "hello"
	byteStr2 := []byte(str3)
	fmt.Println(string(byteStr2))
	byteStr2[3] = 'k'
	fmt.Println(string(byteStr2))

	nan := math.NaN()
	fmt.Println(nan == nan, nan != nan, nan, 1 != 1)

	ab := big.NewInt(0)
	ab1 := big.NewInt(1)
	fmt.Println(ab.Add(ab1, ab))


	var sb strings.Builder
	for i := 0; i < 1000; i++ {
		sb.WriteString("a")
	}
	fmt.Println(sb.String())
}
