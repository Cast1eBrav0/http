.PHONY: build
build:
	go build -v ./
.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build -o goweb-linux -v ./ 
.DEFAULT_GOAL := build