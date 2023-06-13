package main

import (
	"fmt"
	"gorm.io/driver/mysql"

	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PERSON struct {
	Id   int
	Name string
	Age  int
	Tall int
}

func (u PERSON) TableName() string {
	return "person"
}

func main() {
	db, err := gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3307)/mytest"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.Model(&PERSON{}).Where("id=?", 1).Update("name", "zzzz")
}
