# Grpc 





## What is it ?  

It's a framework that can be used for remote interface calls .<br/> 
In grpc , a client application can directly call a method on a server application on a different machine as if it were a local object, making it easier for you to create distributed applications and services. 

![](https://grpc.io/img/landing-2.svg)




## How to use it  in golang  


First you had to ready some tools 

1. Go 



2. Protocol Buffer 

    It's protoc compile tools. it is used for generating pb.go file accroding to `proto` file you defined. 

    you can go  [this](https://github.com/protocolbuffers/protobuf/releases) . and select the suitable version for your OS. 

3. Go plugin for the protocol compiler  

    run the following command to install plugin 

    ```
    go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

    ``` 

    update your PATH , so that protoc compiler can find the plugins:
    ```
    export PATH="$PATH:$(go env GOPATH)/bin"    
    ```

Second , you can define your own `proto file`

```proto 
//hello.proto 
syntax = "proto3";

package api;
option go_package = ".;api";

message responseData{
    int32 A_DATA = 1;
    double B_DATA = 2;

}

message request{
  int32 Id = 1;
}

service AccessDB{
   rpc getDetectInfos(request) returns(responseData){}
}
```

And , Now use `protoc` tools to generate `pb.go` file . Use the following command. `hello.pb.go` will be generated. 

```
protoc --go_out=plugins=grpc:. *.proto
```

Now , you can write your client or server code. Please refer offical examples about `grpc` [link](https://grpc.io/docs/languages/go/quickstart/)

