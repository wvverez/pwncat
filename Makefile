BINARY=pwncat
BUILD_DIR=build

all: build

build:
	go build -o $(BUILD_DIR)/$(BINARY) cmd/pwncat/main.go

clean:
	rm -rf $(BUILD_DIR)

test:
	go test -v ./...

install:
	go install cmd/pwncat/main.go

release:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-linux-amd64 cmd/pwncat/main.go
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-windows-amd64.exe cmd/pwncat/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY)-mac-amd64 cmd/pwncat/main.go
