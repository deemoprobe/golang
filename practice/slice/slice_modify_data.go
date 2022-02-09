package main

// 验证切片值传递

import (
    "fmt"
)

func main() {
    mySlice := []int{10, 20, 30, 40, 50}
    
    // 改变值不要使用这种方式
    // 这种方式只是改变了临时内存里的值
    // 并没有改变实际切片内存内的值
    for _, value := range mySlice {
        value *= 2
    }
    fmt.Printf("mySlice %+v\n", mySlice)

    // 通过索引遍历并改变实际切片的值
    for index, _ := range mySlice {
        mySlice[index] *= 2
    }
    fmt.Printf("mySlice %+v\n", mySlice)
}
