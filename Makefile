
build:
	go build ./cmd/image-build-daemon
.PHONY: deps

check:
	go test ./...
.PHONY: deps

deps:
	glide init -v --skip-tests
.PHONY: deps
