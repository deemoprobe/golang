// +build ignore

package main

import "fmt"

// PersonInfo struct
type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main()  {
	// 声明字典map string是键的类型,PersonInfo存放值
	var personDB map[string] PersonInfo
	
	// // 创建该map string键类型,PersonInfo值类型
	// personDB = make(map[string] PersonInfo)
	// // 给map赋值
	// personDB["123"] = PersonInfo{"3150701334", "xiyunpeng", "Shanghai, China"}
	// personDB["134"] = PersonInfo{"11111110", "Lucy", "Home"}

	// 创建并初始化赋值,相当于上面注释的两步
	personDB = map[string] PersonInfo {
		"123": PersonInfo{"3150701334", "xiyunpeng", "Shanghai, China"},
		"134": PersonInfo{"11111110", "Lucy", "Home"},
	}

	// 查找键为123的信息
	person, ok := personDB["123"]
	if ok {
		fmt.Println("Found person", person.Name, "with key 123 and his ID is", person.ID)
	} else {
		fmt.Println("Sorry! There is no person with key 123.")
	}
	// 删除键为123的信息,即从map中删除数据
	delete(personDB, "123")
	// 再次查找键为123的信息
	person1, ok := personDB["123"]
	println("Now delete person 123. And check again!")
	if ok {
		fmt.Println("Found person", person1.Name, "with key 123 and his ID is", person1.ID)
	} else {
		fmt.Println("Sorry! There is no person with key 123 or you have deleted it.")
	}
}