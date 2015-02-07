## Prepare
#### Install Virtual Box
- Download: [Downloads](https://www.virtualbox.org/wiki/Downloads)
- Install

#### Install docker
Run the following command.
    
    $ sudo yum install docker # CentOs 7+
[Official Doc](https://docs.docker.com/installation/centos/)

#### Install Golang
- Install golang  [Official Doc](http://golang.org/doc/install#tarball)
- set environmental variables 

    GOROOT=#your golang installed path

    GOPATH=%dir%/tester-on-docker/
    
    PATH=%PATH%;%GOROOT%/bin



## Run test
#### start boot2docker
Run the following command.
```
    $ sudo service docker start
```

#### run test

Run the following command.
```
    go test -test.timeout=20m linux_demo
```