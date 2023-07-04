package learn_map

import "fmt"

func DeepCopy() {
	m := make(map[string]string, 0)
	m["name"] = "zhou"
	m["age"] = "24"

	type person struct {
		M map[string]string
	}

	p1 := person{
		M: m,
	}
	fmt.Println(p1) // {map[age:24 name:zhou]}
	m["mm"] = "mm"
	fmt.Println(p1) // {map[age:24 mm:mm name:zhou]}
}
