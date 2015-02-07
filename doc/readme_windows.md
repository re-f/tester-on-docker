## 准备
#### 安装 Virtual Box
- 下载: [官方下载地址](https://www.virtualbox.org/wiki/Downloads)
- Install

#### 安装docker
- 下载： [官方下载地址](https://github.com/boot2docker/windows-installer/releases/tag/v1.3.2)
- 安装： [官方文档](https://docs.docker.com/installation/windows/)

#### 配置环境
- 安装golang，正确设置环境变量（GOROOT，GOPATH，PATH）


    GOROOT=#golang 的安装目录

    GOPATH=c:/users/tester-on-docker/
    
    PATH=%PATH%;%GOROOT%/bin

**注意：c:/users是boot2docker虚拟机建立后设置的默认共享文件夹目录，如果你修改boot2docker的享文件夹目录，那GOPATH的值应为 _你设置的共享文件夹目录/tester-on-docker_**

    [golang 安装参考](http://golang.org/doc/install#windows)

- 配置golang交叉编译
安装mingw([参考](https://github.com/golang/go/wiki/WindowsBuild))
在命令行中执行：
```bat
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	cd %GOROOT%/src
	make.bat
```
[交叉配置参考](https://code.google.com/p/go-wiki/wiki/WindowsCrossCompiling)

## 使用
#### 部署
将tester-on-docker 整个目录复制到c:/users(或者是你设置的boot2docker共享目录)

#### 启动boot2docker并获取boot2docker 的ip
在git bash 中执行
```bash
    $ boot2docker init   #启动boot2docker
    $ boot2docker start   #启动boot2docker
    $ boot2docker ip      #查看boot2docker ip地址
```

####配置文件 
- 修改tester-on-docker/src/windows_demo/test-on-docker.conf 文件，将ssh.ip改为boot2docker ip地址

####运行测试
- 在tester-on-docker/src/demo/windows目录下执行
   
```
    go test -test.timeout=20m
```