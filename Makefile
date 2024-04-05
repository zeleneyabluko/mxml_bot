VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

build: format
	go get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${shell dpkg --print-architecture} go build -v -o mxml_bot -ldflags "-X="github.com/zeleneyabluko/mxml_bot/cmd.appVersion=${VERSION}

clean:
	rm -rf mxml_bot