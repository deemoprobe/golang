package main

/*go vet main.go
* # command-line-arguments
* ./main.go:15:5: Println call has possible formatting directive %d
* ./main.go:16:5: Println call has possible formatting directive %s
*/

import (
    "fmt"
)

func main() {
    name := "test govet"
    fmt.Println("%d\n", name)
    fmt.Println("%s\n", name, name)
}
