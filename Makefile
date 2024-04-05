VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
APP=$(shell basename $(shell git remote get-url origin))
TARGETOS=linux
REGISTRY=svitlanadmytrenko
TARGETARCH=arm64

format:
	gofmt -s -w ./

lint:
	golint

get:
	go test -v

get:
	go get

build: format get
	
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o mxml_bot -ldflags "-X="github.com/zeleneyabluko/mxml_bot/cmd.appVersion=${VERSION}
	
image:
	docker build . -t $(REGISTRY)/$(APP):$(VERSION)-$(TARGETARCH)

push:
	docker push $(REGISTRY)/$(APP):$(VERSION)-$(TARGETARCH)


clean:
	rm -rf mxml_bot
	docker rmi $(REGISTRY)/$(APP):$(VERSION)-$(TARGETARCH)