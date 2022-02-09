package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("Ark", "World", "Hi")
	flag.Parse()
	fmt.Println("Os arg is:", os.Args)
	fmt.Println("Input parameter is:", *name)
	fullstring := fmt.Sprintf("Hello %s from golang.\n", *name)
	fmt.Println(fullstring)
}
