// Package simplemath calc
package simplemath

import "math"

// Sqrt calc
func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}