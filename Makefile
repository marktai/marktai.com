export GOPATH := $(shell pwd)
default: build

init:
<<<<<<< HEAD
	@rm -f bin/main
	@cd src/main && go get
=======
	rm -rf bin/main
	cd src/main && go get
>>>>>>> 753d79d4d7b505be029c75c0c2526dc43fc3fe2a

update: 
	ssh ubuntu@$(INSTANCE) cd /opt/www/marktai.com && git pull && make server &

<<<<<<< HEAD

build: init
	@go build -o bin/main src/main/main.go 

run: build
	bin/main

=======
server: init
	go build -o bin/main src/main/main.go 
	bin/main >> log.txt
>>>>>>> 753d79d4d7b505be029c75c0c2526dc43fc3fe2a
