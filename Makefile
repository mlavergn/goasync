###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := run

.PHONY: test

run:
	cd demo; go run demo.go

mod:
	go mod init async

build:
	go build main.go

test:
	go test -v -count=1 ./...

format:
	go fmt

lint:
	go vet

mac:
	GOOS=darwin GOARCH=arm64 go build main.go

linux:
	GOOS=linux GOARCH=amd64 go build main.go

github:
	open "https://github.com/mlavergn/goasync"

release:
	zip -r goasync.zip LICENSE README.md Makefile cmd src
	hub release create -m "${VERSION} - Go Async" -a goasync.zip -t master "v${VERSION}"
	open "https://github.com/mlavergn/goasync/releases"

st:
	open -a SourceTree .