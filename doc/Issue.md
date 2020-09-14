# Issues

报错1:`main redeclared in this block`  
分析:同一个目录下所有源文件同名称的包只能声明一次  
方案:

- 新建一个目录，将含有相同包的源码文件分开
- 或者在新创建的源码文件开头加入`// +build ignore`构建约束，解决冲突（记得空一行）

报错2:`don't use ALL_CAPS in Go names; use CamelCase`  
分析:不能使用下划线命名法，使用驼峰命名法  
方案:变量用小驼峰方式标识(例如:myBook),方法类等用大驼峰方式标识(例如:MyBook)

报错3:`exported function Xxx should have comment or be unexported`  
分析:外部可见程序结构体、变量、函数都需要注释  
方案:添加注释即可

报错4:`var statJsonByte should be statJSONByte`,`var taskId should be taskID`  
分析:通用名词要求**大写**  
方案:例如以下

- iD/Id -> ID
- Http -> HTTP
- Json -> JSON
- Url -> URL
- Ip -> IP
- Sql -> SQL

报错5:`don't use an underscore in package name`,`don't use MixedCaps in package name; xxXxx should be xxxxx`  
分析:包命名统一小写不使用驼峰和下划线  
方案:包名全部小写即可

报错6:`comment on exported type Repo should be of the form "Repo ..."`  
分析:注释第一个单词要求是注释程序主体的名称  
方案:比如以下实例注释

```go
// Package example add and sub
package example

// Add a+b
func Add(a int, b int) int {
    return a + b
}

// Sub a-b
func Sub(a int, b int) int {
    return a - b
}
```

报错7:`type name will be used as user.UserModel by other packages, and that stutters; consider calling this Model`  
分析:外部可见程序实体不建议再加包名前缀  
方案:同上

报错8:`if block ends with a return statement, so drop this else and outdent its block`  
分析:if语句包含return时，后续代码不能包含在else里面  
方案:同上

报错9:`should replace errors.New(fmt.Sprintf(...)) with fmt.Errorf(...)`  
分析:errors.New(fmt.Sprintf(…)) 建议写成 fmt.Errorf(…)  
方案:同上

报错10:`receiver name should be a reflection of its identity; don't use generic names such as "this" or "self"`  
分析:receiver名称不能为this或self  
方案:同上

报错11:`error var SampleError should have name of the form ErrSample`  
分析:错误变量命名需以 Err/err 开头  
方案:不要以Err/err开头命名变量

报错12:`should replace num += 1 with num++`,`should replace num -= 1 with num--`  
分析:a+=1应该改成a++，a-=1应该改成a–  
方案:同上
