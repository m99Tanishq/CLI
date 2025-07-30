# GLM CLI

A powerful command-line interface for interacting with GLM (General Language Model) APIs, featuring AI-powered code analysis, file manipulation, and codebase memory management.

## 🚀 Features

- **🤖 AI Chat**: Interactive chat sessions with GLM models
- **📁 File Management**: Create, read, write, and search files
- **🔧 Code Analysis**: AI-powered code analysis and fixing
- **🧠 Memory System**: Index and query entire codebases
- **⚙️ Configuration**: Manage API keys and settings
- **📚 History**: Track and manage chat conversations

## 📦 Installation

### Option 1: Quick Install Script (Recommended)

```bash
# Clone the repository
git clone https://github.com/m99Tanishq/glm-cli.git
cd glm-cli

# Run the installation script
./scripts/install.sh
```

The script will:
- ✅ Check if Go is installed
- ✅ Install glm-cli globally
- ✅ Verify the installation
- ✅ Provide PATH setup instructions

### Option 2: Go Install (Global)

```bash
# Install directly from GitHub
go install github.com/m99Tanishq/glm-cli@latest

# Make sure ~/go/bin is in your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Option 3: Direct Download

1. **Download the latest release** for your platform:
   - [Linux x64](https://github.com/m99Tanishq/glm-cli/releases/latest/download/glm-cli-linux-amd64)
   - [macOS x64](https://github.com/m99Tanishq/glm-cli/releases/latest/download/glm-cli-darwin-amd64)
   - [Windows x64](https://github.com/m99Tanishq/glm-cli/releases/latest/download/glm-cli-windows-amd64.exe)

2. **Make it executable** (Linux/macOS):
   ```bash
   chmod +x glm-cli
   ```

3. **Move to your PATH**:
   ```bash
   # Linux/macOS (system-wide)
   sudo mv glm-cli /usr/local/bin/
   
   # Or user-local installation
   mkdir -p ~/.local/bin
   mv glm-cli ~/.local/bin/
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
   source ~/.bashrc
   ```

### Option 4: Build from Source

```bash
git clone https://github.com/m99Tanishq/glm-cli.git
cd glm-cli

# Build and install
make install-check

# Or build manually
go build -o glm-cli .
sudo mv glm-cli /usr/local/bin/
```

### 🔧 PATH Setup

If glm-cli is not found after installation, add the appropriate directory to your PATH:

**Note:** The installation script automatically detects your shell and provides the correct PATH setup instructions.

**For Go install:**
```bash
# Add to ~/.bashrc, ~/.zshrc, or ~/.profile
export PATH="$PATH:$(go env GOPATH)/bin"
```

**For user-local installation:**
```bash
# Add to ~/.bashrc, ~/.zshrc, or ~/.profile
export PATH="$PATH:$HOME/.local/bin"
```

**Verify installation:**
```bash
glm-cli --version
```

### 🗑️ Uninstallation

```bash
# If installed via the install script
./scripts/uninstall.sh

# If installed via go install
rm $(go env GOPATH)/bin/glm-cli

# If installed manually
sudo rm /usr/local/bin/glm-cli
# or
rm ~/.local/bin/glm-cli
```

## 🔧 Setup

1. **Get a Hugging Face API Token**:
   - Go to [Hugging Face](https://huggingface.co/)
   - Create an account and get your access token
   - Your token should start with `hf_`

2. **Configure the CLI**:
   ```bash
   ./glm-cli config --set api_key=YOUR_HF_TOKEN
   ./glm-cli config --set model=zai-org/GLM-4.5:novita
   ```

3. **Verify Configuration**:
   ```bash
   ./glm-cli config --list
   ```

## 🎯 Quick Start

### Chat with AI
```bash
# Start interactive chat
./glm-cli chat

# Chat with specific model
./glm-cli chat --model zai-org/GLM-4.5:novita
```

### File Operations
```bash
# List files
./glm-cli files list

# Read a file
./glm-cli files read myfile.txt

# Create a file
./glm-cli files create newfile.txt

# Write content
./glm-cli files write newfile.txt "Hello World"
```

### Code Analysis
```bash
# Analyze code
./glm-cli code analyze myfile.go

# Fix code issues
./glm-cli code fix myfile.go

# Code review
./glm-cli code review myfile.go
```

### Memory System
```bash
# Index a codebase
./glm-cli memory index .

# Query the codebase
./glm-cli memory query "What is the main function?"

# List indexed data
./glm-cli memory list

# Analyze codebase
./glm-cli memory analyze
```

## 📋 Commands

### Chat Commands
- `chat` - Start interactive chat session
- `chat --model <model>` - Use specific model
- `chat --stream` - Enable streaming responses

### File Commands
- `files list [dir]` - List files in directory
- `files read <file>` - Read file contents
- `files write <file> <content>` - Write to file
- `files create <file>` - Create new file
- `files search <dir> <pattern>` - Search files

### Code Commands
- `code analyze <file>` - Analyze code with AI
- `code fix <file>` - Fix code issues with AI
- `code review <file>` - Code review with AI

### Memory Commands
- `memory index [path]` - Index codebase
- `memory query <query>` - Query indexed codebase
- `memory list` - List indexed data
- `memory analyze [path]` - Analyze codebase
- `memory clear` - Clear indexed data

### Configuration Commands
- `config --set key=value` - Set configuration
- `config --get key` - Get configuration value
- `config --list` - List all configuration

## 🔑 Configuration

The CLI stores configuration in `~/.glm-cli/config.json`:

```json
{
  "api_key": "hf_your_token_here",
  "model": "zai-org/GLM-4.5:novita",
  "base_url": "https://router.huggingface.co/v1",
  "max_history": 100
}
```

### Environment Variables
You can also use environment variables:
```bash
export GLM_API_KEY=hf_your_token_here
export GLM_MODEL=zai-org/GLM-4.5:novita
```

## 🏗️ Project Structure

```
glm-cli/
├── cmd/           # Command implementations
│   ├── chat.go    # Chat functionality
│   ├── code.go    # Code analysis
│   ├── config.go  # Configuration
│   ├── files.go   # File operations
│   ├── history.go # History management
│   ├── memory.go  # Memory system
│   └── root.go    # Root command
├── internal/      # Internal packages
│   ├── api/       # API client
│   ├── config/    # Configuration
│   ├── history/   # History management
│   └── memory/    # Memory system
├── pkg/           # Public packages
│   ├── models/    # Data models
│   └── utils/     # Utilities
├── main.go        # Entry point
└── go.mod         # Go module
```

## 🛠️ Development

### Prerequisites
- Go 1.21 or later
- Git

### Build
```bash
git clone https://github.com/m99Tanishq/glm-cli.git
cd glm-cli
go build -o glm-cli .
```

### Run Tests
```bash
go test ./...
```

### Cross-Platform Build
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o glm-cli-linux-amd64 .

# macOS
GOOS=darwin GOARCH=amd64 go build -o glm-cli-darwin-amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o glm-cli-windows-amd64.exe .
```

## 🚀 Deployment & CI/CD

This project uses GitHub Actions for automated builds, testing, and releases.

### Workflows

#### 1. **CI Pipeline** (`.github/workflows/ci.yml`)
- **Triggers**: Push to `main`/`develop`, Pull Requests
- **Jobs**:
  - **Test**: Runs unit tests, race detection, and build verification
  - **Lint**: Code linting with golangci-lint
  - **Security**: Vulnerability scanning with govulncheck

#### 2. **Development Build** (`.github/workflows/dev-build.yml`)
- **Triggers**: Push to `main` branch
- **Actions**:
  - Builds for all platforms (Linux, macOS, Windows - AMD64/ARM64)
  - Generates SHA256 checksums
  - Uploads artifacts for 30-day retention

#### 3. **Release** (`.github/workflows/release.yml`)
- **Triggers**: Push tags starting with `v*` (e.g., `v1.0.0`)
- **Actions**:
  - Builds optimized binaries for all platforms
  - Injects version information into binaries
  - Generates SHA256 checksums for security
  - Creates GitHub release with all artifacts
  - Auto-generates release notes

### Creating a Release

1. **Update version** in your code
2. **Create and push a tag**:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. **Monitor the release workflow** in GitHub Actions
4. **Download artifacts** from the generated release

### Local Development

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

The CLI includes version information that's injected during build:

```bash
# Check version
./glm-cli version

# Build with custom version
go build -ldflags="-X github.com/m99Tanishq/glm-cli/cmd.Version=v1.0.0" -o glm-cli .
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

- **Issues**: [GitHub Issues](https://github.com/m99Tanishq/glm-cli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/m99Tanishq/glm-cli/discussions)

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- Powered by [GLM-4.5](https://huggingface.co/zai-org/GLM-4.5) model
- API access via [Hugging Face](https://huggingface.co/)

---

**Made with ❤️ for the developer community** 