package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

// 将user结构体字段返回为string
func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.name = "deemoprobe"
	m.age = 26

	println(m.ToString())
}
