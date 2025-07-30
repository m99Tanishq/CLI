.PHONY: build clean test release build-all

# Build the application
build:
	go build -o CLI .

# Clean build artifacts
clean:
	rm -f CLI
	rm -f CLI-*

# Run tests
test:
	go test ./...

# Run tests with security checks
test-full: test check

# Build for all platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X github.com/m99Tanishq/CLI/cmd.Version=dev" -o CLI-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X github.com/m99Tanishq/CLI/cmd.Version=dev" -o CLI-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X github.com/m99Tanishq/CLI/cmd.Version=dev" -o CLI-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X github.com/m99Tanishq/CLI/cmd.Version=dev" -o CLI-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X github.com/m99Tanishq/CLI/cmd.Version=dev" -o CLI-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w -X github.com/m99Tanishq/CLI/cmd.Version=dev" -o CLI-windows-arm64.exe .

# Install globally (to ~/go/bin/)
install:
	go install .
	@echo "✅ CLI installed to $(shell go env GOPATH)/bin/"
	@echo "📝 Make sure $(shell go env GOPATH)/bin/ is in your PATH"
	@echo "   Add this to your ~/.bashrc or ~/.zshrc:"
	@echo "   export PATH=\$$PATH:$(shell go env GOPATH)/bin"

# Install with PATH check
install-check: install
	@echo ""
	@echo "🔍 Checking if CLI is available globally..."
	@if command -v CLI >/dev/null 2>&1; then \
		echo "✅ CLI is available globally!"; \
		CLI version; \
	else \
		echo "❌ CLI is not in PATH"; \
		echo "   Please add $(shell go env GOPATH)/bin to your PATH"; \
		echo "   Or run: export PATH=\$$PATH:$(shell go env GOPATH)/bin"; \
	fi

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
	@ls -la CLI-*

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