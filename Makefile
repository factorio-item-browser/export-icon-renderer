.PHONY: help build clean fix

BUILD_DIR   := build
COMMAND_DIR := cmd
VERSION     ?= dev

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: $(COMMAND_DIR)/* ## Builds the binary of the project.
	for DIRECTORY in $^; do \
		NAME=`basename $$DIRECTORY`; \
		echo "Building $$NAME..."; \
		GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -v -mod vendor -i -o $(BUILD_DIR)/$${NAME}_$(VERSION) $$DIRECTORY/$$NAME.go; \
	done

clean: ## Cleans the project.
	go clean -v
	rm -rf $(BUILD_DIR)

fix: ## Fixes the coding style of the project.
	go fmt ./cmd/... ./pkg/...

vendor: ## Builds or updates the vendor directory
	go mod tidy -v
	go mod vendor -v