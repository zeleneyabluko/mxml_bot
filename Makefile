VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

format:
	gofmt -s -w ./

build:
	go build -v -o mxml_bot -ldflags "-X="github.com/zeleneyabluko/mxml_bot/cmd.appVersion=${VERSION}