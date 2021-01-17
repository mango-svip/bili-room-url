GOARCH=amd64
CGO_ENABLED=0
GOBUILD=go build
OUTPUT=./bin
APP_NAME=bili-room-url

usage:
	@echo "make init"
	@echo "make build"
	@echo "make clean"

init:
	mkdir -p bin
	go mod tidy
clean:
	rm -rf  ./bin
build:
	CGO_ENABLED=${CGO_ENABLED} GOOS=darwin GOARCH=${GOARCH} ${GOBUILD} -o ${OUTPUT}/${APP_NAME} .
build-windows:
	mkdir -p ./bin/windows
	CGO_ENABLED=1 GOOS=windows GOARCH=${GOARCH} ${GOBUILD} -o ${OUTPUT}/windows/${APP_NAME}.exe  .
build-linux:
	mkdir -p ./bin/linux
	CGO_ENABLED=${CGO_ENABLED} GOOS=linux GOARCH=${GOARCH} ${GOBUILD} -o ${OUTPUT}/linux/${APP_NAME}   .

build-all: clean build build-osx build-linux