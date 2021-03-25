# GPC

Golang Process Command

## Environment Setting

1. Go Lang Install(v1.16)
2. [ProtoBuffer Install](https://github.com/protocolbuffers/protobuf)
3. Dependency Setting

### Go Lang Install(v1.16)

```shell
// windows
go env -w GO111MODULE=auto
set GOPATH=E:\SourceCode\gpc
set GOBIN=E:\SourceCode\gpc\bin
set Path=%PATH%;%GOPATH%/bin
```

```shell
// mac os
go env -w GO111MODULE=auto
go env -w GOPATH=/Users/imgeunseog/Documents/Github/gpc
go env -w GOBIN=/Users/imgeunseog/Documents/Github/gpc/bin
export PATH=$PATH:$GOPATH/bin
```

## Build

```shell
go build -o ./dist/gpc gpc.go
```
