## 准备
### 安装docker
- 下载：

	 [官方下载地址](https://github.com/boot2docker/windows-installer/releases/tag/v1.3.2)
- 安装：
	
	[官方文档](https://docs.docker.com/installation/windows/)

- 配置映射目录：
 windows版docker1.3默认将c:/Users 在boot2docker上映射到  /c/Users目录下，无法修改其他目录,如果需要自定义映射目录解决方式参照：[boot2docker together with VirtualBox Guest Additions](https://medium.com/boot2docker-lightweight-linux-for-docker/boot2docker-together-with-virtualbox-guest-additions-da1e3ab2465c)

### 配置环境
- 安装golang，正确设置环境变量（GOROOT，GOPATH，PATH）
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
### 启动boot2docker
在git bash 中执行
```bash
    $ boot2docker start   #启动boot2docker
    $ boot2docker ip      #查看boot2docker ip地址
```

###配置文件
- 将test-on-docker_windows.conf 拷贝到映射目录下，并改名为test-on-docker.conf
- 修改配置文件，将ssh.ip改为boot2docker ip地址

###运行测试
- 在映射目录下执行go test docker

  __注意：测试的运行目录必须在映射目录下__