GO := go
GOBUILD := $(GO) build
package := github.com/cbrnrd/ipgen

default: build

build:
	mkdir -p bin
	$(GOBUILD) -o bin/$(BINARY_NAME) -v

clean:
	rm -f bin/$(BINARY_NAME)

benchmark:
	GOMAXPROCS=1 $(GO) test -bench=. -benchmem -benchtime=10s -cpuprofile $(package)