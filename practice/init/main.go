package main

// 引入两个本地包，这两个包需要先创建
// 内容如下：
// $GOROOT/src/a/init.go
/*
package a

import (
    "fmt"

    _ "b"
)

func init() {
    fmt.Println("init from a")
}
*/
// $GOROOT/src/b/init.go
/*
package b

import (
    "fmt"
)

func init() {
    fmt.Println("init from b")
}
*/

import (
    "fmt"
    _ "a"
    _ "b"
)

func init() {
    fmt.Println("main init")
}

func main() {
    fmt.Println("main func")
}

/*
结果：
# go run main.go 
init from b
init from a
main init
main func
可见启动顺序：b--a--init--main
*/
