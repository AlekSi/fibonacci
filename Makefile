# using grep here due to https://github.com/golang/go/issues/11659
DIRS = $(shell go list ./... | grep -v vendor/)

all: test

install:
	go install -v $(DIRS)

check: install
	go tool vet -all -shadow $(shell ls -d */ | grep -v vendor/)
	- errcheck $(DIRS)
	- golint ./... | grep -v vendor/

test: install
	go test -i -v $(DIRS)
	go test -v $(DIRS)

run: install
	go run main.go -debug
