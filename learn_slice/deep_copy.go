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
