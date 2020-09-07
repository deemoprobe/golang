# Golang

## 介绍

Golang是Google开发的开源编程语言,主要目标是"兼具Python 等动态语言的开发速度和C/C++等编译型语言的性能与安全性".Go语言已经⼴泛应用于人工智能、云计算开发、容器虚拟化、⼤数据开发、数据分析及科学计算、运维开发、爬虫开发、游戏开发等领域.

## 特征

- 语法简单
- 并发模型
    Goroutine处理并发单元,goroutine类似于线程,但并非线程.可以将goroutine理解为一种虚拟线程.Go语言运行时会参与调度 goroutine,并将goroutine合理地分配到每个CPU中,最大限度地使用CPU性能.而且开启一个goroutine的消耗非常小(约2KB内存).
    Goroutine特点：
  - `goroutine`具有可增长的分段堆栈。这意味着它们只在需要时才会使用更多内存.
  - `goroutine`的启动时间比线程快.
  - `goroutine`原生支持利用channel安全地进行通信.
  - `goroutine`共享数据结构时无需使用互斥锁.
- 内存分配
- 垃圾回收
- 静态链接
- 标准库
- 工具链

## 方向

- 服务端开发
- 分布式系统，微服务
- 网络编程
- 区块链开发
- 内存KV数据库，例如boltDB、levelDB
- 云平台

## 内容

### 工程结构

GOROOT=go安装目录(安装go后自动生成)
GOPATH=项目目录(可以创建多个,加入环境变量即可)
GOBIN=可执行文件路径(一般不需要设置,这样编译生成的可执行文件会自动保存在GPATH\bin下)
工程项目结构：
![20200907153330](https://deemoprobe.oss-cn-shanghai.aliyuncs.com/images/20200907153330.png)

```shell
cd gopath/src
go install myapp1 // 编译,在bin目录下生成可执行文件
```
