# using grep here due to https://github.com/golang/go/issues/11659
DIRS = $(shell go list ./... | grep -v vendor/)

all: test

install:
	go install -v $(DIRS)

install-race:
	go install -v -race $(DIRS)

check: install
	go tool vet -all -shadow $(shell ls -d */ | grep -v vendor/)
	- errcheck $(DIRS)
	- golint ./... | grep -v vendor/

test: install
	go test -i -v $(DIRS)
	go test -v $(DIRS)

test-race: install-race
	go test -i -v -race $(DIRS)
	env GORACE="halt_on_error=1" go test -v -race $(DIRS)

run: install
	go run main.go -debug

run-race: install-race
	go run -race main.go -debug
