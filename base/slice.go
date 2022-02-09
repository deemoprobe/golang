package main

import (
	"fmt"
)

func main() {
	// 创建容量为5的切片
	x := make([]int, 0, 5)

	for i := 0; i < 10; i++ {
		// 追加数据，容量超标会自动增大容量
		x = append(x, i)
	}

	fmt.Println(x)
}
