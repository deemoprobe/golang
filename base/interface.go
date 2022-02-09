package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

// 接口是一组方法的集合
type Printer interface {
	Print()
}

func main() {
	var u user
	u.name = "deemoprobe"
	u.age = 26

	var p Printer = u
	p.Print()
}
