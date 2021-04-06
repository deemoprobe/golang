# Go 国内加速镜像

## 说明

众所周知,国内网络访问国外资源经常会出现不稳定的情况. Go 生态系统中有着许多中国 Gopher 们无法获取的模块,比如最著名的 golang.org/x/...并且在中国大陆从 GitHub 获取模块的速度也有点慢.

因此设置 CDN 加速代理就很有必要了,以下是几个速度不错的提供者:

* 七牛:Goproxy 中国 <https://goproxy.cn>
* 阿里: <https://mirrors.aliyun.com/goproxy/>
* 官方: 全球 CDN 加速 <https://goproxy.io>

## 设置代理

### 类 Unix

在 Linux 或 macOS 上面,需要运行下面命令（或者,可以把以下命令写到 .bashrc 或 .bash_profile 文件中）:

```shell
# 启用 Go Modules 功能
go env -w GO111MODULE=on

# 配置 GOPROXY 环境变量,以下三选一

# 1. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn,direct

# 2. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 3. 官方
go env -w  GOPROXY=https://goproxy.io,direct
确认一下:

$ go env | grep GOPROXY
GOPROXY="https://goproxy.cn"
```

### Windows

在 Windows 上,git bash 中运行上面相同的命令即可
