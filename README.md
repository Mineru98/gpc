# GPC

Golang Process Command

## Environment Setting

1. Go Lang Install(v1.16)
2. [ProtoBuffer Install](https://github.com/protocolbuffers/protobuf)
3. Dependency Setting


### Go Lang Install(v1.16)
```shell
go env -w GO111MODULE=auto
go env -w GOPATH=/Users/imgeunseog/Documents/Github/gpc
go env -w GOBIN=/Users/imgeunseog/Documents/Github/gpc/bin
export PATH=$PATH:$GOPATH/bin
```

### Dependency Setting
```shell
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/akamensky/argparse
go get -u github.com/jaypipes/ghw
```

## Build
```shell
go build -o ./dist/gpc gpc.go
```