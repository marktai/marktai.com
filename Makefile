export GOPATH := $(shell pwd)

update: 
	ssh ubuntu@$(INSTANCE) cd /opt/www/marktai.com && git pull 
server: 
	go run src/main/main.go
