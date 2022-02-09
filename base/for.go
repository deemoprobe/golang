package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	for i := 10; i >= 5; i-- {
		println(i)
	}

	x := []int{100, 200, 300}

	for i, n := range x {
		println(i, ":", n)
	}
}
