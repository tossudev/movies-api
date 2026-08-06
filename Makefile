init:
	git config core.hooksPath .githooks

build:
	go build

test:
	go test -v
