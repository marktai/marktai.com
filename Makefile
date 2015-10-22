export GOPATH := $(shell pwd)
default: build

init:
	@rm -f bin/main
	@cd src/main && go get

update: 
	ssh ubuntu@$(INSTANCE) cd /opt/www/marktai.com && git pull && make server


build: init
	@go build -o bin/main src/main/main.go 

run: build
	bin/main

