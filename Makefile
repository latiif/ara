VERSION = v0.7
COMMITID = $(shell git rev-parse HEAD)
build:
	go mod tidy
	rm -f ara
	go build -ldflags "-X github.com/latiif/ara/cmd.VERSION=$(VERSION) -X github.com/latiif/ara/cmd.COMMITID=$(COMMITID)"  -o ara
