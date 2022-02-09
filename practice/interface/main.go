package main

/*
接口：一组方法的集合，go语言多态的一种实现
*/

import "fmt"

type Info interface {
    getInfo() string
}

type Human struct {
    firstName, lastName string
}

type Cat struct {
    kind string
    name string
}

type Dog struct {
    name string
    age string
}

func (h Human) getInfo() string {
    return "---\nHuman:\nname:" + h.firstName + " " + h.lastName
}

func (c *Cat) getInfo() string {
    return fmt.Sprintf("---\nCat:\nname:%s, kind:%s", c.name, c.kind)
}

func (d *Dog) getInfo() string {
    return "---\nDog:\nname:" + d.name + ", age:" + d.age
}

func main() {
    // 定义interface类型的切片
    inter := []Info{}
    
    h := new(Human)
    h.firstName = "Hua"
    h.lastName = "Li"
    inter = append(inter, h)
    
    c := new(Cat)
    c.kind = "blue cat"
    c.name = "kitty"
    inter = append(inter, c)
    
    for _, f := range inter {
        fmt.Println(f.getInfo())
    }

    d := Dog{}
    d.name = "little hei"
    d.age = "2"
    fmt.Println(d.getInfo())
}

