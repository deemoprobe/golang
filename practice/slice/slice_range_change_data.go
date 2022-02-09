package main

// 1.修改结构体数组中特定名称
// 2.修改结构体数组中所有名为jack为lucy

import (
    "fmt"
)

type User struct {
    name string
    age int
}

func main() {
    mySlice := []User{{"deemoprobe", 26}, {"deemo", 10}, {"jack", 22}, {"jack", 19}}
    for index,_ := range mySlice {
        fmt.Printf("%s %d ", mySlice[index].name, mySlice[index].age)
    }

    // 修改特定下标的数据
    for index,_ := range mySlice {
        if index == 1 {
            mySlice[index].name = "lihua"
        }
    }

    // 修改所有匹配的数据
    for index,_ := range mySlice {
        if mySlice[index].name == "jack" {
            mySlice[index].name = "lucy"
        }
    }

    fmt.Println()
    for index,_ := range mySlice {
        fmt.Printf("%s %d ", mySlice[index].name, mySlice[index].age)
    }
}
