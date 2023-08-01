package learn_slice

import "fmt"

func DeepCopy() {
	s := make([]string, 0)
	s = append(s, "zhou")
	s = append(s, "24")

	type person struct {
		S []string
	}

	p1 := person{
		S: s,
	}
	fmt.Println(p1) // [zhou 24]
	s = append(s, "mm")
	fmt.Println(p1) // [zhou 24]
}

func DeepCopyPtr() {
	s := make([]string, 0)
	s = append(s, "zhou")
	s = append(s, "24")

	type person struct {
		S *[]string
	}

	p1 := person{
		S: &s,
	}
	fmt.Println(*p1.S) // [zhou 24]
	s = append(s, "mm")
	fmt.Println(*p1.S) // [zhou 24 mm]
}

func Append() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a[2:5:9]
	fmt.Printf("&s %p &s[0] %p, %v %v %v\n", &s, &s[0], s, len(s), cap(s)) //&s 0xc00009c0c0 &s[0] 0xc0000a4060, [3 4 5] 3 7
	//s = append(s, 11, 21, 31, 41)
	appendV(s)
	fmt.Println(s, len(s), cap(s)) //[3 4 5] 3 7
	fmt.Println(a)                 //[1 2 3 4 5 11 21 31 9 10]

	s1 := s[0:2]
	fmt.Println(s1, len(s1), cap(s1)) //[3 4] 2 7
	s2 := s1[0:2]
	fmt.Println(s2, len(s2), cap(s2)) //[3 4] 2 7
	s2[0] = 10
	fmt.Println(s, s1, s2, a) //[10 4 5] [10 4] [10 4] [1 2 10 4 5 11 21 31 9 10]
}

func appendV(s []int) {
	s = append(s, 11, 21, 31)
	fmt.Printf("in &s %p &s[0] %p, %v %v %v\n", &s, &s[0], s, len(s), cap(s)) //in &s 0xc00009c0f0 &s[0] 0xc0000a4060, [3 4 5 11 21 31] 6 7
}
