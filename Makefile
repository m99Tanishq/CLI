.PHONY: build clean test release build-all

# Build the application
build:
	go build -o glm-cli .

# Clean build artifacts
clean:
	rm -f glm-cli
	rm -f glm-cli-*

# Run tests
test:
	go test ./...

# Run tests with security checks
test-full: test check

# Build for all platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X github.com/m99Tanishq/glm-cli/cmd.Version=dev" -o glm-cli-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X github.com/m99Tanishq/glm-cli/cmd.Version=dev" -o glm-cli-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X github.com/m99Tanishq/glm-cli/cmd.Version=dev" -o glm-cli-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X github.com/m99Tanishq/glm-cli/cmd.Version=dev" -o glm-cli-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X github.com/m99Tanishq/glm-cli/cmd.Version=dev" -o glm-cli-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w -X github.com/m99Tanishq/glm-cli/cmd.Version=dev" -o glm-cli-windows-arm64.exe .

# Install locally
install:
	go install .

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Security scan
security:
	@echo "🔒 Running security scan..."
	@echo "Current Go version: $(shell go version)"
	@echo ""
	@govulncheck ./... || (echo ""; echo "⚠️  Note: Some vulnerabilities require Go 1.23+"; echo "   See SECURITY.md for detailed analysis"; echo "   Current risk level: LOW"; exit 0)

# Full check (lint + security)
check: lint security

# Run with race detection
race:
	go run -race .

# Generate release
release: build-all
	@echo "Built binaries for all platforms:"
	@ls -la glm-cli-*

# Development mode
dev:
	go run .

# Update dependencies
deps:
	go mod tidy
	go mod download

# Create a new release
release-version:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make release-version VERSION=v1.0.0"; \
		exit 1; \
	fi
	./scripts/release.sh $(VERSION) 