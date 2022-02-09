package main

import (
    "fmt"
)

// 回调函数

func main() {
    x, y := 10, 2
    a := FunCallBack(x, y, Add)
    b := FunCallBack(x, y, Sub)
    c := FunCallBack(x, y, Multi)
    d := FunCallBack(x, y, Div)
    fmt.Println("x=", x)
    fmt.Println("y=", y)
    fmt.Println("x+y=", a)
    fmt.Println("x-y=", b)
    fmt.Println("x*y=", c)
    fmt.Println("x/y=", d)
}

// 定义加减乘除
func Add(a, b int) int {
    return a + b
}

func Sub(a, b int) int {
    return a - b
}

func Multi(a, b int) int {
    return a * b
}

func Div(a, b int) int {
    return a / b
}

type Callback func(x, y int) int

// 定义回调函数
func FunCallBack(x, y int, f Callback) int {
    return f(x, y)
}
/*上面几行（type开始）等价于下面
func FunCallBack(x, y int, f (func(x, y int) int)) int{
    return f(x, y)
}
*/

