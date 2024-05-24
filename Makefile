# Variables
SRC_DIR := .

# Go commands
GO := go
GO_TEST := $(GO) test
GO_VET := $(GO) vet
GO_FMT := $(GO) fmt
GO_MOD := $(GO) mod
GO_LINT := golangci-lint run


# Run tests
test:
	@echo "Running tests..."
	$(GO_TEST) -v ./...

# Format the code
fmt:
	@echo "Formatting code..."
	$(GO_FMT) ./...

# Vet the code
vet:
	@echo "Vetting code..."
	$(GO_VET) ./...

# Lint the code
lint:
	@echo "Linting code..."
	$(GO_LINT) ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO_MOD) tidy

profile:
	go test -coverprofile=coverage.out ./...

html:
	go tool cover -html=coverage.out

	
# Default target
all: test lint fmt vet profile html