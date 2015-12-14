export GOPATH := $(shell pwd)
default: build

init:
	rm -f bin/server bin/main
	@cd src/main && go get

build: init
	go build -o bin/server src/main/main.go 

run: build
	bin/server>log.txt 2>&1 &
