## 准备
#### 安装 Virtual Box
- 下载: [官方下载地址](https://www.virtualbox.org/wiki/Downloads)
- 安装

#### 安装docker
在命令行中执行
    
    $ sudo yum install docker # CentOs 7+
[官方文档](https://docs.docker.com/installation/centos/)

#### 安装Golang
- 安装golang，[golang 安装参考](http://golang.org/doc/install#tarball)
- 设置环境变量（GOROOT，GOPATH，PATH）
  
    GOROOT=#golang 的安装目录

    GOPATH=%dir%/tester-on-docker/
    
    PATH=%PATH%;%GOROOT%/bin

## 使用
#### 启动docker
在命令行中执行
```
    $ sudo service docker start
```

####运行测试
在命令行中执行   
```
    go test -test.timeout=20m linux_demo
```