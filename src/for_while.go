// +build ignore

package main

func main() {
	x := 0

	// * golang  只有三种流控制语句,分别是if/for/switch
	// ! 这里只是相当于while控制流
	for x < 5 {
		println(x)
		x++
	}
}
