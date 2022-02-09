package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

type manager struct {
	// 匿名嵌入其他类型
	user
	title string
}

func main() {
	var m manager

	m.name = "deemoprobe"
	m.age = 26
	m.title = "Engineer"

	fmt.Println(m)
}
