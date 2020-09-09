// +build ignore

package main

import "fmt"

// Modify array value for now
func Modify(array [5]int) {
	array[0] = 10
	fmt.Println("In Modify(), array values:", array)
}

func main()  {
	array := [5]int{1,2,3,4,5}
	Modify(array)
	// 函数内的操作不影响主函数中的数组输出
	fmt.Println("In main(), array values:", array)
}