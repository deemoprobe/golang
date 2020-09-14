// +build ignore

package main

import "fmt"

func main() {
	var x int32
	var s = "hello go."
	// 函数内部可以省略var
	y := 64

	// 定义的参数如果未使用会报错
	fmt.Println(x, s, y)
}
