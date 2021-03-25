hello :
	@echo "Hello World"

run :
	@go run cmd/gpc.go

gpc-cc:
	@echo "Cross Complie"
	set GOOS=darwin& set GOARCH=arm64& go build -o bin/gpc-darwin-arm64 cmd/gpc.go
	set GOOS=darwin& set GOARCH=amd64& go build -o bin/gpc-darwin-amd64 cmd/gpc.go
	set GOOS=linux& set GOARCH=386& go build -o bin/gpc-linux-386 cmd/gpc.go
	set GOOS=linux& set GOARCH=arm64& go build -o bin/gpc-linux-amd64 cmd/gpc.go
	set GOOS=linux& set GOARCH=amd64& go build -o bin/gpc-linux-amd64 cmd/gpc.go
	set GOOS=windows& set GOARCH=386& go build -o bin/gpc-windows-386.exe cmd/gpc.go
	set GOOS=windows& set GOARCH=amd64& go build -o bin/gpc-windows-amd64.exe cmd/gpc.go

master-cc:
	@echo "Cross Complie"
	set GOOS=darwin& set GOARCH=arm64& go build -o bin/master-darwin-arm64 cmd/master.go
	set GOOS=darwin& set GOARCH=amd64& go build -o bin/master-darwin-amd64 cmd/master.go
	set GOOS=linux& set GOARCH=386& go build -o bin/master-linux-386 cmd/master.go
	set GOOS=linux& set GOARCH=arm64& go build -o bin/master-linux-amd64 cmd/master.go
	set GOOS=linux& set GOARCH=amd64& go build -o bin/master-linux-amd64 cmd/master.go
	set GOOS=windows& set GOARCH=386& go build -o bin/master-windows-386.exe cmd/master.go
	set GOOS=windows& set GOARCH=amd64& go build -o bin/master-windows-amd64.exe cmd/master.go

slave-cc:
	@echo "Cross Complie"
	set GOOS=darwin& set GOARCH=arm64& go build -o bin/slave-darwin-arm64 cmd/slave.go
	set GOOS=darwin& set GOARCH=amd64& go build -o bin/slave-darwin-amd64 cmd/slave.go
	set GOOS=linux& set GOARCH=386& go build -o bin/slave-linux-386 cmd/slave.go
	set GOOS=linux& set GOARCH=arm64& go build -o bin/slave-linux-amd64 cmd/slave.go
	set GOOS=linux& set GOARCH=amd64& go build -o bin/slave-linux-amd64 cmd/slave.go
	set GOOS=windows& set GOARCH=386& go build -o bin/slave-windows-386.exe cmd/slave.go
	set GOOS=windows& set GOARCH=amd64& go build -o bin/slave-windows-amd64.exe cmd/slave.go

gpc-cc-unix:
	@echo "Cross Complie"
	GOOS=darwin GOARCH=amd64 go build -o bin/gpc-darwin-amd64 cmd/gpc.go
	GOOS=darwin GOARCH=arm64 go build -o bin/gpc-darwin-arm64 cmd/gpc.go
	GOOS=linux GOARCH=386 go build -o bin/gpc-linux-386 cmd/gpc.go
	GOOS=linux GOARCH=amd64 go build -o bin/gpc-linux-amd64 cmd/gpc.go
	GOOS=linux GOARCH=arm64 go build -o bin/gpc-linux-amd64 cmd/gpc.go
	GOOS=windows GOARCH=386 go build -o bin/gpc-windows-386.exe cmd/gpc.go
	GOOS=windows GOARCH=amd64 go build -o bin/gpc-windows-amd64.exe cmd/gpc.go

master-cc-unix:
	@echo "Cross Complie"
	GOOS=darwin GOARCH=amd64 go build -o bin/master-darwin-amd64 cmd/master.go
	GOOS=darwin GOARCH=arm64 go build -o bin/master-darwin-arm64 cmd/master.go
	GOOS=linux GOARCH=386 go build -o bin/master-linux-386 cmd/master.go
	GOOS=linux GOARCH=amd64 go build -o bin/master-linux-amd64 cmd/master.go
	GOOS=linux GOARCH=arm64 go build -o bin/master-linux-amd64 cmd/master.go
	GOOS=windows GOARCH=386 go build -o bin/master-windows-386.exe cmd/master.go
	GOOS=windows GOARCH=amd64 go build -o bin/master-windows-amd64.exe cmd/master.go

slave-cc-unix:
	@echo "Cross Complie"
	GOOS=darwin GOARCH=amd64 go build -o bin/slave-darwin-amd64 cmd/slave.go
	GOOS=darwin GOARCH=arm64 go build -o bin/slave-darwin-arm64 cmd/slave.go
	GOOS=linux GOARCH=386 go build -o bin/slave-linux-386 cmd/slave.go
	GOOS=linux GOARCH=amd64 go build -o bin/slave-linux-amd64 cmd/slave.go
	GOOS=linux GOARCH=arm64 go build -o bin/slave-linux-amd64 cmd/slave.go
	GOOS=windows GOARCH=386 go build -o bin/slave-windows-386.exe cmd/slave.go
	GOOS=windows GOARCH=amd64 go build -o bin/slave-windows-amd64.exe cmd/slave.go