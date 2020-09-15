# GoTips

1. 安装go_tools

    配置go代理解决go-tools因被墙无法安装的问题<https://learnku.com/go/wikis/38122>

2. 在src中创建主程序目录和包程序目录,分别放置对应源代码文件,编译时在bin目录下执行`go build 主程序名`,此时在  
bin目录下会生成对应的可执行文件,以及在pkg目录下会生成对应的编译包文件,例如下面:

    ```shell
    bin> go build example
    bin> ls
    bin> example.exe
    bin> cd ../pkg
    bin> ls
    pkg> example.a
    ```
