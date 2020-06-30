
build:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o GO_algos cmd/main.go
.PHONY: build

build-mac:
	CGO_ENABLED=0 \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o GO_algos cmd/main.go
.PHONY: build-mac

build-windows:
	CGO_ENABLED=0 \
	GOOS=windows \
	GOARCH=amd64 \
	go build -o GO_algos.exe cmd/main.go
.PHONY: build-windows

test:
	@echo "Starting Unit Tests"
	go test -v ./... -tags=unit
.PHONY: unit-test