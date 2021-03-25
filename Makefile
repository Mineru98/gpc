hello :
	@echo "Hello World"

run :
	@go run cmd/gpc.go

cc:
	@echo "Cross Complie"
	set GOOS=darwin& set GOARCH=arm64& go build -o bin/gpc-darwin-arm64 cmd/gpc.go
	set GOOS=darwin& set GOARCH=amd64& go build -o bin/gpc-darwin-amd64 cmd/gpc.go
	set GOOS=linux& set GOARCH=386& go build -o bin/gpc-linux-386 cmd/gpc.go
	set GOOS=linux& set GOARCH=arm64& go build -o bin/gpc-linux-amd64 cmd/gpc.go
	set GOOS=linux& set GOARCH=amd64& go build -o bin/gpc-linux-amd64 cmd/gpc.go
	set GOOS=windows& set GOARCH=386& go build -o bin/gpc-windows-386.exe cmd/gpc.go
	set GOOS=windows& set GOARCH=amd64& go build -o bin/gpc-windows-amd64.exe cmd/gpc.go

cclinux-arm:
	@echo "Cross Complie"
	GOOS=darwin GOARCH=amd64 go build -o bin/gpc-darwin-amd64 cmd/gpc.go
	GOOS=darwin GOARCH=arm64 go build -o bin/gpc-darwin-arm64 cmd/gpc.go
	GOOS=linux GOARCH=386 go build -o bin/gpc-linux-386 cmd/gpc.go
	GOOS=linux GOARCH=amd64 go build -o bin/gpc-linux-amd64 cmd/gpc.go
	GOOS=linux GOARCH=arm64 go build -o bin/gpc-linux-amd64 cmd/gpc.go
	GOOS=windows GOARCH=386 go build -o bin/gpc-windows-386.exe cmd/gpc.go
	GOOS=windows GOARCH=amd64 go build -o bin/gpc-windows-amd64.exe cmd/gpc.go