# gpc

Golang Process Command

## Developement Environment Setting

1. Go Lang Install
2. [ProtoBuffer Install](https://github.com/protocolbuffers/protobuf)
3. gRPC setting

```shell
go env -w GO111MODULE=auto
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

```shell
go env -w GO111MODULE=on
go env -w GOPATH=/Users/imgeunseog/Documents/Github/gpc
go env -w GOBIN=/Users/imgeunseog/Documents/Github/gpc/bin
export PATH=$PATH:$GOPATH/bin
```