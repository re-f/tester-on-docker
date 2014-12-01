## 准备工作

### 安装docker
- 下载：
	 [官方下载地址](https://github.com/boot2docker/windows-installer/releases/tag/v1.3.2)
- 安装：
	
	[官方文档](https://docs.docker.com/installation/windows/)

- 配置映射目录：
 windows版docker1.3+默认将c:/Users 在boot2docker上映射到  /c/Users目录下，无法修改其他目录
 解决方式参照：[boot2docker together with VirtualBox Guest Additions](https://medium.com/boot2docker-lightweight-linux-for-docker/boot2docker-together-with-virtualbox-guest-additions-da1e3ab2465c),可自定义映射目录

 __注意：测试的执行文件或测试源代码必须放在映射目录下__

### 配置galang交叉编译环境

## 使用
###配置test-on-docker.conf
具体参数含义，参照test-on-docker.conf.proto，将配置文件放在
示例参照test-on-docker_windows.conf




