# Makefile to build Golang service

# Service
SERVICE_DIR := service
SERVICE_BINARY := service

build:
	cd $(SERVICE_DIR)/ ; go build -o $(SERVICE_BINARY)

clean: ## Removes binaries
	rm $(SERVICE_DIR)/$(SERVICE_BINARY) || true

run: clean build
	cd $(SERVICE_DIR)/ ; ./$(SERVICE_BINARY)

lint: ## Run linting for each module
	golangci-lint run ./...

test:
	go test -v ./...