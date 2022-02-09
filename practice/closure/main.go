package main

/* 
闭包：由函数和相关引用环境的组合实体
函数：调用外部函数
环境：外部函数的环境，本实例中是变量x
f1() f2()就是闭包函数
*/

import (
    "fmt"
)

func main() {
    // 第一次赋值
    x := 100
    f1 := Result1(&x)
    f2 := Result2(x)

    // 1
    fmt.Println("Answer1:")
    f1()
    f1()
    f2()
    f2()

    // 2
    // 二次赋值
    x = 200
    fmt.Println("Answer2:")
    f1()
    f1()
    f2()
    f2()

    // 3
    fmt.Println("Answer3:")
    Result1(&x)()
    Result1(&x)()
    Result2(x)()
    Result2(x)()
}

// 外部函数，通过指针方式传值，每次调用自增1
func Result1(x *int) func() {
  return func() {
    *x = *x + 1
    fmt.Printf("Result1 = %d\n", *x)
  }
}

// 外部函数，直接进行值传递，每次调用自增1
func Result2(x int) func() {
  return func() {
    x = x + 1
    fmt.Printf("Result2 = %d\n", x)
  }
}

