// +build ignore

package main

func main() {
	x := 4

	// * for if 模拟while true
	for {
		println(x)
		x--

		if x < 0 {
			break
		}
	}
}
