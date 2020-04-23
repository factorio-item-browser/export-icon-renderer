.PHONY: help build clean fix test vendor

BUILD_DIR   := build
COMMAND_DIR := cmd
PACKAGE     := github.com/factorio-item-browser/export-icon-renderer
VERSION     ?= dev

BUILD_FLAGS := -v -i -mod vendor -ldflags "-s -w -X $(PACKAGE)/pkg/env.Version=$(VERSION)"

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
	#go test ./cmd/... ./pkg/...
	go vet ./cmd/... ./pkg/...

vendor: ## Builds or updates the vendor directory
	go mod tidy -v
	go mod vendor -v