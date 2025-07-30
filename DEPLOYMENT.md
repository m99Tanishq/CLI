# Deployment Guide

This document outlines the deployment and CI/CD setup for the GLM CLI project.

## 🚀 GitHub Workflows

### 1. CI Pipeline (`ci.yml`)
**Purpose**: Continuous Integration for code quality and testing
**Triggers**: 
- Push to `main` or `develop` branches
- Pull requests to `main` branch

**Jobs**:
- **Test**: Unit tests, race detection, build verification
- **Lint**: Code linting with golangci-lint
- **Security**: Vulnerability scanning with govulncheck

### 2. Development Build (`dev-build.yml`)
**Purpose**: Build artifacts for development testing
**Triggers**: Push to `main` branch

**Actions**:
- Builds for all platforms (Linux, macOS, Windows - AMD64/ARM64)
- Generates SHA256 checksums for security
- Uploads artifacts for 30-day retention

### 3. Release (`release.yml`)
**Purpose**: Create official releases
**Triggers**: Push tags starting with `v*` (e.g., `v1.0.0`)

**Actions**:
- Builds optimized binaries for all platforms
- Injects version information into binaries
- Generates SHA256 checksums
- Creates GitHub release with all artifacts
- Auto-generates release notes

## 📋 Release Process

### Automated Release
1. **Create a tag**:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **Or use the release script**:
   ```bash
   ./scripts/release.sh v1.0.0
   ```

3. **Or use Make**:
   ```bash
   make release-version VERSION=v1.0.0
   ```

### Manual Release Steps
1. Ensure you're on the `main` branch
2. Pull latest changes: `git pull origin main`
3. Run tests: `make test`
4. Build for all platforms: `make build-all`
5. Create and push tag: `git tag v1.0.0 && git push origin v1.0.0`
6. Monitor GitHub Actions workflow
7. Download release artifacts

## 🔧 Local Development

### Build Commands
```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Format code
make fmt

# Lint code
make lint

# Clean build artifacts
make clean
```

### Version Management
```bash
# Check current version
./glm-cli version

# Build with custom version
go build -ldflags="-X github.com/m99Tanishq/glm-cli/cmd.Version=v1.0.0" -o glm-cli .
```

## 📦 Artifacts

### Release Artifacts
Each release includes:
- `glm-cli-linux-amd64` - Linux x64 binary
- `glm-cli-linux-arm64` - Linux ARM64 binary
- `glm-cli-darwin-amd64` - macOS x64 binary
- `glm-cli-darwin-arm64` - macOS ARM64 binary
- `glm-cli-windows-amd64.exe` - Windows x64 binary
- `glm-cli-windows-arm64.exe` - Windows ARM64 binary
- SHA256 checksums for all binaries

### Development Artifacts
Development builds are available as GitHub Actions artifacts for 30 days.

## 🔒 Security

- All binaries are built with security flags (`-ldflags="-s -w"`)
- SHA256 checksums are generated for integrity verification
- Security scanning is performed on every CI run
- Dependencies are regularly updated and scanned

## 📊 Monitoring

- **Workflow Status**: https://github.com/m99Tanishq/glm-cli/actions
- **Releases**: https://github.com/m99Tanishq/glm-cli/releases
- **Issues**: https://github.com/m99Tanishq/glm-cli/issues

## 🛠️ Troubleshooting

### Common Issues

1. **Build fails on GitHub Actions**
   - Check Go version compatibility
   - Verify all dependencies are available
   - Check for syntax errors

2. **Release not created**
   - Ensure tag format is correct (`v*`)
   - Check GitHub Actions permissions
   - Verify workflow file syntax

3. **Version not showing correctly**
   - Check ldflags syntax in workflow
   - Verify variable name matches code
   - Rebuild with correct flags

### Debug Commands
```bash
# Check workflow syntax
yamllint .github/workflows/*.yml

# Test build locally
make build-all

# Verify version injection
./glm-cli version
```

## 📈 Future Enhancements

- [ ] Add Docker container builds
- [ ] Implement automated dependency updates
- [ ] Add performance benchmarking
- [ ] Include code coverage reporting
- [ ] Add automated changelog generation 