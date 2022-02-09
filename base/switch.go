package main

func main() {
	x := 100

	switch {
	case x > 10:
		println(x)
	case x < 10:
		println(-x)
	default:
		println(0)
	}
}
