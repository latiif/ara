VERSION = v0.5.1
COMMITID = $(shell git rev-parse HEAD)
build:
	rm -f ara
	go build -ldflags "-X github.com/latiif/ara/cmd.VERSION=$(VERSION) -X github.com/latiif/ara/cmd.COMMITID=$(COMMITID)"  -o ara