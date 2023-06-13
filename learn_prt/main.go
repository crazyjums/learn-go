package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) SetName(s string) {
	c.Name = s
}

func (c Cat) GetName() string {
	return c.Name
}

func (c *Cat) SetNameV2(s string) {
	c.Name = s
}

func (c *Cat) GetNameV2() string {
	return c.Name
}

func main() {
	var c1, c2 Cat
	c1.SetName("cat1")                     // 等价于 (&c1).SetName("cat1")
	c2.SetNameV2("cat2")                   // 等价于 (&c2).SetNameV2("cat2")
	fmt.Printf("c1: %s\n", c1.GetName())   // ""
	fmt.Printf("c2: %s\n", c2.GetNameV2()) // "cat2"

	var c3, c4 *Cat
	c3 = &Cat{}
	c4 = &Cat{}
	c3.SetName("cat3")
	c4.SetNameV2("cat4")
	fmt.Printf("c3: %s\n", c3.GetName())   // ""
	fmt.Printf("c4: %s\n", c4.GetNameV2()) // "cat4"
}
