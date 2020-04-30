.PHONY: dev build install test release clean

all: dev

dev: build
	@./static -v

build: clean
	@go build \
		-tags "netgo static_build" -installsuffix netgo \
		-ldflags "-w -X $(shell go list).Version=$(VERSION) -X $(shell go list).Commit=$(COMMIT)" \
		.

install: build
	@go install

test:
	@go test ./...

release:
	@./tools/release.sh

clean:
	@git clean -d -f -x
