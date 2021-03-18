# gpc

Golang Process Command

## Developement Environment Setting

1. Go Lang Install
2. [ProtoBuffer Install](https://github.com/protocolbuffers/protobuf)
3. gRPC setting

```shell
go env -w GO111MODULE=on
go env -w GOPATH=/Users/imgeunseog/Documents/Github/gpc
go env -w GOBIN=/Users/imgeunseog/Documents/Github/gpc/bin
export PATH=$PATH:$GOPATH/bin
```

```shell
go env -w GO111MODULE=auto
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/akamensky/argparse
go get -u github.com/jaypipes/ghw
```

## Build
```shell
go build -o ./dist/gpc gpc.go
```