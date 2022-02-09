package main

import (
    "fmt"
)

func main() {
    
    // string
    str := "a string value"

    // string pointer
    pointer := &str
    
    // string pointer value
    anotherString := *&str
    
    // print
    fmt.Println(str)
    fmt.Println(pointer)
    fmt.Println(anotherString)
    
    // change string vlaue
    str = "changed string"
    
    // print
    fmt.Println(str)
    fmt.Println(pointer)
    fmt.Println(anotherString)
    
    para := ParameterStruct{Name: "aaa"}
    fmt.Println(para)
    changeParameter(&para, "bbb")
    fmt.Println(para)
    cannotChangeParameter(para, "ccc")
    fmt.Println(para)
}

type ParameterStruct struct {
    Name string
}

func changeParameter(para *ParameterStruct, value string) {
    para.Name = value
}

func cannotChangeParameter(para ParameterStruct, value string) {
    para.Name = value
}

