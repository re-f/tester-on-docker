## Prepare
#### Install Virtual Box
- Download: [Downloads](https://www.virtualbox.org/wiki/Downloads)
- Install

#### Install docker
- Download: [Downloads](https://github.com/boot2docker/windows-installer/releases/tag/v1.3.2)

- Install: [Official Doc](https://docs.docker.com/installation/windows/)

#### Environment
- Install golang and set environmental variables 

    GOROOT=#your golang installed path

    GOPATH=c:/users/tester-on-docker/
    
    PATH=%PATH%;%GOROOT%/bin

**attention：c:/users is boot2docker default share folder，if you change it GAPATH value is %share folder%/tester-on-docker**


[golang 安装参考](http://golang.org/doc/install#windows)

- Cross comilation
Install mingw([reference](https://github.com/golang/go/wiki/WindowsBuild))
execute in cmd :
```bat
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	cd %GOROOT%/src
	make.bat
```
[Cross compilation reference](https://code.google.com/p/go-wiki/wiki/WindowsCrossCompiling)

## Run test

#### Deployment
copy tester-on-docker director to boot2docker share folder(default folder is c:/users)    
 
#### start boot2docker & get boot2docker ip
execute on git bash

```bash
    $ boot2docker init    #initiate boot2docker
    $ boot2docker start   #start boot2docker
    $ boot2docker ip      #show boot2docker ip addresss
```

####config file
update tester-on-docker/src/windows_demo/test-on-docker.conf，set ssh.ip to boot2docker ip address

####run test
execute on cmd
```bash
    cd c:/users/tester-on-docker/src/demo/windows
    go test -test.timeout=20m docker
```