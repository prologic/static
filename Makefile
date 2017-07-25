.PHONY: dev build clean

all: dev

dev: build
	./static

build: clean
	go get ./...
	go build .

test:
	go test ./...

clean:
	rm -rf static
