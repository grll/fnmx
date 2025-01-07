.PHONY: all clean

BINARY_NAME=fnmx
VERSION=1.0.0

all: clean build-all

clean:
	rm -rf dist/

build-all: clean
	mkdir -p dist
	GOOS=darwin GOARCH=arm64 go build -o dist/${BINARY_NAME}-darwin-arm64
	GOOS=darwin GOARCH=amd64 go build -o dist/${BINARY_NAME}-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o dist/${BINARY_NAME}-windows-amd64.exe
	GOOS=windows GOARCH=arm64 go build -o dist/${BINARY_NAME}-windows-arm64.exe

# For local testing
build:
	go build -o ${BINARY_NAME}