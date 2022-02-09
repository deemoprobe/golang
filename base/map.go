package main

import (
	"fmt"
)

func main() {
	// 创建字典类型对象
	m := make(map[string]int)

	m["a"] = 1

	// ok-idiom模式，判定操作是否成功
	// map["b"]不存在，所以会失败
	x, ok := m["b"]
	fmt.Println(x, ok)

	y, ok := m["a"]
	fmt.Println(y, ok)

	delete(m, "a")
}
