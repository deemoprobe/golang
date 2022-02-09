package main

/*
selfdefmath包需要提前建好，包含如下四个文件：

# /usr/local/go/src/selfdefmath/add.go

package selfdefmath

// Add calc
func Add(a int, b int) int {
    return a + b
}

# /usr/local/go/src/selfdefmath/add_test.go
package selfdefmath

import "testing"

func TestAdd1(t *testing.T) {
    r := Add(1, 2)
    if r != 3 {
        t.Errorf("Add(1, 2) failed. Got %v, expected 3.", t)
    }
}

# /usr/local/go/src/selfdefmath/sqrt.go 
package selfdefmath

import "math"

// Sqrt calc
func Sqrt(i int) int {
    v := math.Sqrt(float64(i))
    return int(v)
}

# /usr/local/go/src/selfdefmath/sqrt_test.go
package selfdefmath

import "testing"

func TestSqrt1(t * testing.T)  {
    v := Sqrt (16)
    if v !=4 {
        t.Errorf("Sqrt(16) failed. Got %v, expect 4.", v)
    }
}
*/

import (
    "os"
    "fmt"
    "strconv"

    "selfdefmath"
)

// Usage calc
var Usage = func() {
    fmt.Println("Usage: calc command [arg] ... ")
    fmt.Println("\nThe command are:\n\tadd\tAddtion of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main()  {
    args := os.Args[1:]
    if args == nil || len(args) < 2 {
        Usage()
        return
    }

    switch args[0] {
    case "add":
        if len(args) != 3 {
            fmt.Println("Usage: calc add <int1><int2>")
            return
        }
        v1, err1 := strconv.Atoi(args[1])
        v2, err2 := strconv.Atoi(args[2])
        if err1 != nil || err2 != nil {
            fmt.Println("Usage: calc add <int1><int2>")
            return
        }
        ret := selfdefmath.Add(v1,v2)
        fmt.Println("Result: ", ret)
    case "sqrt":
        if len(args) !=2 {
            fmt.Println("Usage: calc aqrt <int>")
            return
        }
        v, err := strconv.Atoi(args[1])
        if err != nil {
            fmt.Println("Usage: calc sqrt <int>")
            return
        }
        ret := selfdefmath.Sqrt(v)
        fmt.Println("Result: ", ret)
    default:
        Usage()
    }
}
