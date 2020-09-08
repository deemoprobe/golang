# LearnGo Notes

## 变量

### 变量声明

go中使用var来声明变量,将类型放在变量后说明

```go
var v1 int      // int
var v2 string   // 字符
var v3 [10]int  // 数组
var v4 []int    // 数组切片
var v5 struct { // 结构体
    f int
}
var v6 *int     // 指针
var v7 map[string]int // map, key为string类型,value为int
var v8 func(a, int) int
// 同时声明多个变量
var (
    v1 int
    v2 string
    v3 ...
)
```

### 变量初始化

在初始化时,var关键字可以保留,但不是必要元素

```go
// 下面三种均可以
var v1 int = 10
var v2 = 10 // go会自动解析类型
v3 := 10    // go会自动解析类型
```

### 变量赋值

赋值和初始化不一样,赋值是声明后进行赋值,初始化时声明同时进行赋值,而且go支持多重赋值

```go
// 赋值
var v1 int
v1 = 10
// 多重赋值,比如交换i和j的值
i := 1
j := 2
i, j = j, i
```

### 匿名变量

比如一个函数有多个返回值,但我们只需要取其中一个值,不要取所有的值,这时就可以用匿名变量

```go
fun GetName() (firstName, lastName, nickName string) {
    return "May" "Chan" "Chi Mar"
}
// 如果只需要获取nickName,调用可以直接写成
_, _, nickName := GetName()
```

## 常量

### 常量定义

常量定义时可以限定常量的类型,但也不是必要条件

```go
const Pi float64 = 3.1415926
const zero = 0.0    // 无类型浮点常量
const (
    size int = 1024
    df = -1     // 无类型整型常量
)
const u, v float32 = 0, 3 // u = 0.0, v = 3.0 常量的多重赋值
const a, b, c = 1, 2, "test" // a = 1, b = 2, c = "test"
const mask = 1 << 3 // 也可以是常量表达式
```

### 预定义常量

go中有三个预定义常量: true, false, iota  
iota可以认为是一个可被编译器修改的常量,在const出现时被重置为0,然后在下一个const出现之前iota出现一次就自增1

```go
const (
    a = iota    // a = 0
    b = iota    // b = 1
)
const (
    c = iota    // c = 0
    d = iota    // d = 1
)

// const 中常量的赋值一样的话可以简写为下面
const (
    a = iota    // a = 0
    b           // b = 1
)
const (
    c = iota    // c = 0
    d           // d = 1
)
```

## 枚举

枚举指的是定义一系列相关的常量合集

```go
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    numberOfDays
)
```

## 类型

go内置以下基础类型

* 布尔
