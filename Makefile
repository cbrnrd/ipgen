GO := go
GOBUILD := $(GO) build
package := github.com/cbrnrd/ipgen
VERSION ?= $(shell git describe --tags --always --dirty)

default: build

build:
	mkdir -p bin
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

clean:
	rm -f bin/$(BINARY_NAME)

bench:
	GOMAXPROCS=1 go test -bench=. -benchmem 

tag:
	@echo "Tagging $(VERSION)"
	@sleep 3
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)