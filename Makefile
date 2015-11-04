export GOPATH := $(shell pwd)
default: build

init:
update: 
	ssh ubuntu@$(INSTANCE) cd /opt/www/marktai.com && git pull && make server &

build: init
	@go build -o bin/main src/main/main.go 

run: build
	bin/main
server: init
	go build -o bin/main src/main/main.go 
	bin/main >> log.txt

