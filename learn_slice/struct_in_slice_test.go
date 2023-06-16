package learn_slice

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
}

func TestStructInSlice(t *testing.T) {
	var persons = make([]*Person, 5, 5)
	for i := range persons {
		persons[i] = &Person{}
	}
	for i, e := range persons {
		fmt.Println(i, e)
	}
	fmt.Println(persons)
}
