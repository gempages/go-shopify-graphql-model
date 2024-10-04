.PHONY: help

TESTDIR?=./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

govet:
	go vet ./...

gotest: govet
gotest: ## Run all test cases or specific cases - example: make gotest TESTDIR=./graphql/model
	go clean -testcache
	go test -cover -race ./...
