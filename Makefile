MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -euo pipefail -c
.DEFAULT_GOAL := all

BIN_DIR ?= $(shell go env GOPATH)/bin
export PATH := $(PATH):$(BIN_DIR)

.PHONY: deps
deps: ## download go modules
	go mod download

.PHONY: vendor
vendor: ## download and vendor dependencies
	go mod vendor

.PHONY: fmt
fmt: lint/check ## ensure consistent code style
	golangci-lint run --fix > /dev/null 2>&1 || true
	go mod tidy

.PHONY: lint/check
lint/check: 
	@if ! golangci-lint --version > /dev/null 2>&1; then \
		echo -e "golangci-lint is not installed: run \`make lint-install\` or install it from https://golangci-lint.run"; \
		exit 1; \
	fi

.PHONY: lint-install
lint-install: ## installs golangci-lint to the go bin dir
	@if ! golangci-lint --version > /dev/null 2>&1; then \
		echo "Installing golangci-lint"; \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.30.0; \
	fi

.PHONY: lint
lint: lint/check ## run golangci-lint
	golangci-lint run

.PHONY: test
test: lint ## run go tests
	go test ./... -race

.PHONY: build
build: ## compile and build artifact
	go build .

.PHONY: all
all: test build

.PHONY: help
help: ## displays this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_\/-]+:.*?## / {printf "\033[34m%-12s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | \
		sort | \
		grep -v '#'
