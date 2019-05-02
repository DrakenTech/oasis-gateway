all:  build test build-gateway

build:
	go build ./...

build-gateway:
	go build github.com/oasislabs/developer-gateway/cmd/gateway

lint:
	go vet ./...
	golangci-lint run

test:
	go test -v -race ./...

test-coverage:
	go test -v -covermode=count -coverprofile=coverage.out ./...
