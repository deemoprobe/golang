// +build ignore

package main

import "fmt"

func main() {
	// 定义数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建数组切片
	// var mySlice []int = myArray[:5]
	var mySlice []int = myArray[:5]
	// 也可以不指定mySlice类型,会自动识别 var mySlice = myArray[:5]
	// 也可以直接赋值 mySlice := myArray[:5]

	fmt.Println("Elements of myArray: ")
	// 遍历数组
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")
	// 遍历数组
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	//fmt.Println()
}
