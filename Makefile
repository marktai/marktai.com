export GOPATH := $(shell pwd)
default: server

init:
	rm -rf bin/main
	cd src/main && go get

update: 
	ssh ubuntu@$(INSTANCE) cd /opt/www/marktai.com && git pull && make server &

server: init
	go build -o bin/main src/main/main.go 
	bin/main > Log.txt
