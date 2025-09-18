# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Binary names
BINARY_NAME=api
BINARY_UNIX=$(BINARY_NAME)_unix

# Build info
VERSION?=$(shell git describe --tags --always --dirty)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# Linker flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: all build clean test coverage deps fmt lint vet help

all: test build ## Run tests and build

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary file
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v ./cmd/api

build-linux: ## Build the binary file for Linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_UNIX) -v ./cmd/api

clean: ## Remove build artifacts
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f coverage.out
	rm -f coverage.html

test: ## Run tests
	$(GOTEST) -v -race -coverprofile=coverage.out ./...

coverage: test ## Run tests and show coverage
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

coverage-func: test ## Show coverage by function
	$(GOCMD) tool cover -func=coverage.out

deps: ## Download dependencies
	$(GOMOD) download
	$(GOMOD) tidy

deps-upgrade: ## Upgrade dependencies
	$(GOGET) -u ./...
	$(GOMOD) tidy

fmt: ## Format code
	$(GOFMT) -s -w .

fmt-check: ## Check if code is formatted
	@test -z $$($(GOFMT) -s -l . | tee /dev/stderr) || (echo "Code not formatted, run 'make fmt'" && exit 1)

lint: ## Run linter
	$(GOLINT) run

vet: ## Run go vet
	$(GOCMD) vet ./...

check: fmt-check vet lint ## Run all checks

run: ## Run the application
	$(GOCMD) run ./cmd/api

run-dev: ## Run the application in development mode
	@echo "Starting development server..."
	@export ENVIRONMENT=development && \
	export LOG_LEVEL=debug && \
	export SERVER_PORT=8080 && \
	$(GOCMD) run ./cmd/api

docker-build: ## Build Docker image
	docker build -t goscaffold:latest .

docker-run: ## Run Docker container
	docker run -p 8080:8080 --env-file .env goscaffold:latest

docker-compose-up: ## Start services with docker-compose
	docker-compose up -d

docker-compose-down: ## Stop services with docker-compose
	docker-compose down

install-tools: ## Install development tools
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint@latest

benchmark: ## Run benchmarks
	$(GOTEST) -bench=. -benchmem ./...

profile: ## Run with profiling
	$(GOCMD) run ./cmd/api -cpuprofile=cpu.prof -memprofile=mem.prof

# Database operations
db-migrate-up: ## Run database migrations up
	@echo "Running database migrations up..."
	# Add your migration command here

db-migrate-down: ## Run database migrations down
	@echo "Running database migrations down..."
	# Add your migration command here

db-reset: ## Reset database
	@echo "Resetting database..."
	# Add your database reset command here

# Release operations
release-patch: ## Create a patch release
	@echo "Creating patch release..."
	@./scripts/release.sh patch

release-minor: ## Create a minor release
	@echo "Creating minor release..."
	@./scripts/release.sh minor

release-major: ## Create a major release
	@echo "Creating major release..."
	@./scripts/release.sh major
