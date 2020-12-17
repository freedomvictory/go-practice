
# go install and configure 

## congigure GOROOT and GOPATH


- GOPATH    always set `$HOME/go`  
- GOROOT    always set golang installtion folder

**reference link**

[SET GOPATH](https://github.com/golang/go/wiki/GOPATH)

[SET GOROOT](https://github.com/golang/go/wiki/MultipleGoRoots)


## configure Chinese golang source 

1. go 1.13 above version

```bash 
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

2. linux 或macos  

```bash
$ export GO111MODULE=on
$ export GOPROXY=https://goproxy.cn
```

## Use module  

1. ***project directroy***

    go project is out of `$GOPATH/src` dir 

2. ***Reference Link***

    [*Go Module Blog*](https://blog.golang.org/using-go-modules)



