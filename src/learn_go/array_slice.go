// +build ignore

package main

import "fmt"

func main()  {
	mySlice := make([]int, 5, 10)

	// 返回数组切片中元素个数
	fmt.Println("len(mySlice)_before:", len(mySlice))
	// 返回数组切片中分配的空间大小
	fmt.Println("cap(mySlice)_before:", cap(mySlice))
	for i, v := range(mySlice) {
		fmt.Println("mySlice_before[", i, "]", v)
	}

	// 向数组中加入三个元素
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("len(mySlice)_after:", len(mySlice))
	fmt.Println("cap(mySlice)_after:", cap(mySlice))
	for i, v := range(mySlice) {
		fmt.Println("mySlice_after[", i, "]", v)
	}

	// 向数组中加入另一个数组的元素,注意省略号
	mySlice2 := []int{4, 5, 6}
	mySlice = append(mySlice, mySlice2...)
	// 上面语句实际上相当于
	// mySlice = append(mySlice, 4, 5)
	fmt.Println("len(mySlice)_after_after:", len(mySlice))
	fmt.Println("cap(mySlice)_after_after:", cap(mySlice))
	for i, v := range(mySlice) {
		fmt.Println("mySlice_after_after[", i, "]", v)
	}

	// 输出结果发现,cap从10变为20了,这是因为数组切片会自动处理空间不足的问题,当元素空间超出
	// 额定值(即cap()一开始的值)后会自动分配足够大的空间来容纳更多的元素,啥也别说了,牛皮...

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{8, 9}
	
	println("slice1:")
	for _, v := range(slice1) {
		fmt.Print(v, " ")
	}

	// 换行
	fmt.Println()

	println("slice2:")
	for _, v := range(slice2) {
		fmt.Print(v, " ")
	}

	// 换行
	fmt.Println()

	// 复制slice2元素到slice1
	copy(slice1, slice2)
	println("slice_copy_slice2:")
	for _, v := range(slice1) {
		fmt.Print(v, " ")
	}
}