export GOPATH := $(shell pwd)
default: build

init:
	bower install
	cd src/main && go get

clean:
	rm -f bin/server bin/main bin/api

build: init clean
	go build -o bin/api src/main/main.go 

buildBeta: build
	mv bin/api bin/apiBeta

run: build
	@-pkill ^api$
	bin/api>>log.txt 2>&1 &

runBeta: buildBeta
	@-pkill ^apiBeta$
	bin/apiBeta -Port=8082 >>log.txt 2>&1 &
