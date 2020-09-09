package mymath

import "errors"
import "fmt"

// Add a + b
func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers")
		fmt.Println("Error:", err)
		return
	}
	// if err != nil{
	// 	fmt.Println("Error:", err)
	// } else {
	// 	return a + b, nil
	// }
	return a + b, nil
}

