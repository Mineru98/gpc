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
set GOPATH=[PATH]\gpc
set GOBIN=[PATH]\gpc\bin
set Path=%PATH%;%GOPATH%\bin

setx GOPATH [PATH]\gpc
setx GOBIN [PATH]\gpc\bin
setx Path %PATH%;[PATH]\gpc\bin
```

```shell
// mac os
go env -w GO111MODULE=auto
go env -w GOPATH=/Users/imgeunseog/Documents/Github/gpc
go env -w GOBIN=/Users/imgeunseog/Documents/Github/gpc/bin
export PATH=$PATH:$(go env GOPATH)/bin
```

## Dependencis

```shell
go get github.com/akamensky/argparse
go get github.com/jaypipes/ghw
go get github.com/mitchellh/go-ps
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get golang.org/x/sys

go get -u github.com/akamensky/argparse
go get -u github.com/jaypipes/ghw
go get -u github.com/mitchellh/go-ps
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u golang.org/x/sys
```

## Build

```shell
# golang
go build -o ./dist/gpc gpc.go
go build -o ./dist/master master.go
go build -o ./dist/slave slave.go
```

```shell
# proto
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go

protoc --go_out=.  --go-grpc_out=. *.proto
```

## gopls setting

```json
{
    ...,
   "gopls": {
        "build.experimentalWorkspaceModule": true,
        "build.expandWorkspaceToModule": false,
        "build.allowModfileModifications": true
    }
}
````
