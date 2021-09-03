# GPC

Golang Process Command

## Environment Setting

1. Go Lang Install(v1.16)
2. [ProtoBuffer Install](https://github.com/protocolbuffers/protobuf)
3. Dependency Setting

### Go Lang Install(v1.16)

```cmd
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
go env -w GOPATH=~/Documents/Github/gpc
go env -w GOBIN=~/Documents/Github/gpc/bin
export PATH=$PATH:$(go env GOPATH)/bin
```

## Dependencies

```shell
go get golang.org/x/tools/gopls
go get github.com/akamensky/argparse
go get github.com/jaypipes/ghw
go get github.com/mitchellh/go-ps
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get golang.org/x/sys/windows/svc

go get -u golang.org/x/tools/gopls
go get -u github.com/akamensky/argparse
go get -u github.com/jaypipes/ghw
go get -u github.com/mitchellh/go-ps
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get golang.org/x/sys/windows/svc
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
    "go.useLanguageServer": true,
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.snippetSuggestions": "none"
  },
  "[go.mod]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "gopls": {
    "experimentalWorkspaceModule": true,
    "usePlaceholders": true,
    "staticcheck": false
  }
}
```
