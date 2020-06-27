
build:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o algo cmd/main.go
.PHONY: build

build-mac:
	CGO_ENABLED=0 \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o algo cmd/main.go
.PHONY: build-mac

test:
	@echo "Starting Unit Tests"
	go test -v ./... -tags=unit
.PHONY: unit-test