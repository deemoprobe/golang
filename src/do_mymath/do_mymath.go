package main

import "mymath"
import "fmt"

func main()  {
	num,_ := mymath.Add(123, 456)
	fmt.Println(num)
}