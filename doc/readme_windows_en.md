## Prepare
### Install docker
- Download:

	 [Download](https://github.com/boot2docker/windows-installer/releases/tag/v1.3.2)
- Install:
	
	[Officials Doc](https://docs.docker.com/installation/windows/)

- Shared folder:
 The first of the following share names that exists (if any) will be automatically mounted at the location specified:

        Users share at /Users
        /Users share at /Users
        c/Users share at /c/Users
        /c/Users share at /c/Users
        c:/Users share at /c/Users

Customize shard folder reference: [boot2docker together with VirtualBox Guest Additions](https://medium.com/boot2docker-lightweight-linux-for-docker/boot2docker-together-with-virtualbox-guest-additions-da1e3ab2465c)

### Environment
- Install golang and set environmental variables (GOROOT，GOPATH，PATH)
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
[cross compilation reference](https://code.google.com/p/go-wiki/wiki/WindowsCrossCompiling)

## Run test
### start boot2docker
execute on git bash
```bash
    $ boot2docker start   #start boot2docker
    $ boot2docker ip      #show boot2docker ip addresss
```

###config file
- copy test-on-docker_windows.conf to shared folder and rename to test-on-docker.conf
- update test-on-docker.conf，set ssh.ip value boot2docker ip地址

###run test
execute on cmd

```bash
    cd c:/users #enter into shard folder
    go test docker
```

  __ATTENTION:Must run test under shard folder__