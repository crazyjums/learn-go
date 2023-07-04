package learn_struct

import "fmt"

func DeepCopyValue() {
	type data struct {
		name string
		age  int
	}

	type person struct {
		d data
	}

	d := data{
		name: "zhou",
		age:  24,
	}

	p := person{
		d: d,
	}

	fmt.Println(p) // {{zhou 24}}
	d.age = 25
	fmt.Println(p) // {{zhou 24}}
}

func DeepCopyPtr() {
	type data struct {
		name string
		age  int
	}

	type person struct {
		d *data
	}

	d := data{
		name: "zhou",
		age:  24,
	}

	p := person{
		d: &d,
	}

	fmt.Println(*p.d) // {zhou 24}
	d.age = 25
	fmt.Println(*p.d) // {zhou 25}
}
