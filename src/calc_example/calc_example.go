package main

import "fmt"
import "example"

func main() {
	v1 := 1
	v2 := 2
	var x int
	var y int
	x = example.Add(v1, v2)
	fmt.Println("1+2=", x)
	y = example.Sub(v2, v1)
	fmt.Println("2-1=", y)
}
