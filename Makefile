BINARY_NAME=rpc-cmd

setup: ## prepares the project
	go mod download

build-client: ## builds the json rpc cli client
	rm -f cmd/client/main.go
	go run cmd/generator/main.go --client-folder=cmd/client
	rm -f build/linux/$(BINARY_NAME)
	GOOS=linux GOARCH=amd64 go build -o build/linux/$(BINARY_NAME) cmd/client/main.go

build-server: ## builds the json rpc server
	rm -f build/linux/$(BINARY_NAME)-server
	GOOS=linux GOARCH=amd64 go build -o build/linux/$(BINARY_NAME)-server cmd/server/main.go

build-all: build-server build-client  ## builds server and client

ci:
	go mod download
	cd test && go mod download
	cd test && go run test/main.go

# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help