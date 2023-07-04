package learn_struct

import (
	"encoding/json"
	"fmt"
	"log"
)

// Stu 使用 tag 定义了结构体字段与 json 字段的转换关系，Name -> stu_name, ID -> stu_id，忽略 Age 字段。
//很方便地实现了 Go 结构体与不同规范的 json 文本之间的转换。
type Stu struct {
	Name string `json:"stu_name"`
	ID   string `json:"stu_id"`
	Age  int    `json:"-"`
	Next *Stu   `json:"-"`
}

type Person struct {
	Name string
	Age  int8
}

type H struct {
	Car string
}

func (p *Person) setName(name string) {
	p.Name = name
}

func (p *Person) setAge(age int8) {
	p.Age = age
}

type Stu1 struct {
	p Person
}

type Stu2 struct {
	Person //匿名字段
	H
}

func testComponent() {
	stu1 := Stu1{
		Person{
			Name: "tony",
			Age:  19,
		},
	}
	fmt.Println(stu1.p.Name)

	stu2 := Stu2{
		Person{
			Name: "tony",
			Age:  19,
		},
		H{
			Car: "HONDA",
		},
	}
	fmt.Println(stu2.Name, stu2.Car)
}

type Set map[string]struct{}

type Lamp struct{}

func (l Lamp) On() {
	println("On")

}
func (l Lamp) Off() {
	println("Off")
}

func init() {
	log.Println("init...")
}

func (s *Stu) setName(name string) {
	s.Name = name
}

func main() {
	var buf, _ = json.Marshal(Stu{"zhangsan", "1234", 14, &Stu{}})
	fmt.Printf("%s \n", buf)
	set := make(Set)
	set["a"] = struct{}{}
	fmt.Println(set)

	//定义指针类型的结构体，必须先初始化才能对其进行赋值，因为指针的零值是nil
	var s1 *Stu
	s1 = &Stu{}
	s1.setName("定义指针类型的结构体")
	fmt.Println(s1)

	testComponent()
}
