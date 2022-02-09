package main

/*
指针：指针是指向存储变量值内存的地址
指针字节数通常是小于值的字节数，通过
指针操作可以加速内存之间值的传递效率

下面例子：本质上都是值传递，地址均为发生变化
*/

import (
    "fmt"
)

// 定义指针值类型为int的指针变量a, b
// 正确交换函数1
/*
直接对内存映射地址进行操作，直接在相应地址重新赋值
视觉上是让人觉得交换了二者的值
*/
func Swap1(a, b *int) {
    // a指针的值赋给temp
    t := *a
    // b指针的值指向a指针的值
    *a = *b
    // temp值指向b指针的值
    *b = t
}

// 正确交换函数2
/*
异或（0^0=0 0^1=1 1^0=1 1^1=0）实现值交换，反复迭代
一式代入二式：*a = *a ^ (*a ^ *b) = *a ^ *a ^ *b = *b
二式代入三式：*b = (*a ^ *b) ^ *b = *a ^ *b ^ *b = *a
本质上和Swap1一致，都是进行了对应指针地址上值的重新赋值，从而实现交换
直接对内存映射地址进行操作，不需要返回main函数
*/
func Swap2(a, b *int) {
    *b = *a ^ *b
    *a = *a ^ *b
    *b = *a ^ *b
}

// 错误的交换函数1
/*
原因是仅仅在函数内执行了形式参数的交换，没有返回值，main函数接收不到交换后的值
*/
func Swap3(a, b int) {
    t := a
    a = b
    b = t
}

// 错误的交换函数2
/*
原因是仅仅在函数内执行了形式参数的交换，没有返回值，main函数接收不到交换后的指针地址
*/
func Swap4(a, b *int) {
    t := a
    a = b
    b = t 
}

func main() {
    
    x1, y1 := 1, 2
    x2, y2 := 3, 4
    x3, y3 := 5, 6
    x4, y4 := 7, 8

    fmt.Println("Swap1:")
    // 交换前
    fmt.Println("x1, y1=", x1, y1)
    fmt.Println("&x1, &y1=", &x1, &y1)
    // 交换
    Swap1(&x1, &y1)
    // 交换后
    fmt.Println("x1, y1=", x1, y1)
    fmt.Println("&x1, &y1=", &x1, &y1)


    fmt.Println("Swap2:")
    // 交换前
    fmt.Println("x2, y2=", x2, y2)
    fmt.Println("&x2, &y2=", &x2, &y2)
    // 交换
    Swap2(&x2, &y2)
    // 交换后
    fmt.Println("x2, y2=", x2, y2)
    fmt.Println("&x2, &y2=", &x2, &y2)


    fmt.Println("Swap3:")
    // 交换前
    fmt.Println("x3, y3=", x3, y3)
    fmt.Println("&x3, &y3=", &x3, &y3)
    // 交换
    Swap3(x3, y3)
    // 交换后
    fmt.Println("x3, y3=", x3, y3)
    fmt.Println("&x3, &y3=", &x3, &y3)


    fmt.Println("Swap4:")
    // 交换前
    fmt.Println("x4, y4=", x4, y4)
    fmt.Println("&x4, &y4=", &x4, &y4)
    // 交换
    Swap4(&x4, &y4)
    // 交换后
    fmt.Println("x4, y4=", x4, y4)
    fmt.Println("&x4, &y4=", &x4, &y4)
}
