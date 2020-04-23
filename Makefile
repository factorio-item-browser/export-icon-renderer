.PHONY: help build clean fix test test-cover vendor

BUILD_DIR   := build
COMMAND_DIR := cmd
PACKAGE     := github.com/factorio-item-browser/export-icon-renderer
VERSION     ?= dev

LD_FLAGS    := -ldflags "-s -w -X $(PACKAGE)/pkg/env.Version=$(VERSION)"
BUILD_FLAGS := -v -i -mod vendor $(LD_FLAGS)

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: $(COMMAND_DIR)/* ## Builds the binary of the project.
	for DIRECTORY in $^; do \
		NAME=`basename $$DIRECTORY`; \
		echo "Building $$NAME..."; \
		GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$$NAME $$DIRECTORY/$$NAME.go; \
	done

clean: ## Cleans the project.
	go clean -v
	rm -rf $(BUILD_DIR)

fix: ## Fixes the coding style of the project.
	go fmt ./cmd/... ./pkg/...

test: ## Tests the project.
	go test $(LD_FLAGS) ./cmd/... ./pkg/...
	go vet $(LD_FLAGS) ./cmd/... ./pkg/...

test-cover: ## Tests the project and generates the coverage report.
	go test $(LD_FLAGS) -coverprofile=$(BUILD_DIR)/coverage.out ./cmd/... ./pkg/...
	go tool cover -html=$(BUILD_DIR)/coverage.out

vendor: ## Builds or updates the vendor directory
	go mod tidy -v
	go mod vendor -v