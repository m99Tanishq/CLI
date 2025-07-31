# GLM CLI

A powerful command-line interface for interacting with GLM (General Language Model) APIs, featuring AI-powered code analysis, file manipulation, and codebase memory management.

## 🚀 Features

- **🤖 AI Chat**: Interactive chat sessions with GLM models
- **📁 File Management**: Create, read, write, and search files
- **🔧 Code Analysis**: AI-powered code analysis and fixing
- **🧠 Memory System**: Index and query entire codebases
- **⚙️ Configuration**: Manage API keys and settings
- **📚 History**: Track and manage chat conversations

## 💻 System Requirements

- **Operating Systems**: Linux, macOS, Windows
- **Architectures**: AMD64 (x86_64), ARM64 (Apple Silicon, ARM64 Linux)
- **Memory**: 50MB RAM minimum
- **Network**: Internet connection for API access
- **Dependencies**: None (statically linked binary)

## 📦 Installation

### Option 1: Download Latest Release

Download the latest release for your platform:

#### Linux
```bash
# AMD64
wget https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-linux-amd64
chmod +x CLI-linux-amd64
sudo mv CLI-linux-amd64 /usr/local/bin/CLI

# ARM64
wget https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-linux-arm64
chmod +x CLI-linux-arm64
sudo mv CLI-linux-arm64 /usr/local/bin/CLI
```

#### macOS
```bash
# Intel Mac (AMD64)
curl -L -o CLI https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-darwin-amd64
chmod +x CLI
sudo mv CLI /usr/local/bin/

# Apple Silicon (ARM64)
curl -L -o CLI https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-darwin-arm64
chmod +x CLI
sudo mv CLI /usr/local/bin/
```

#### Windows
```powershell
# AMD64
Invoke-WebRequest -Uri "https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-windows-amd64.exe" -OutFile "CLI.exe"
# Move CLI.exe to a directory in your PATH (e.g., C:\Windows\System32)

# ARM64
Invoke-WebRequest -Uri "https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-windows-arm64.exe" -OutFile "CLI.exe"
# Move CLI.exe to a directory in your PATH (e.g., C:\Windows\System32)
```

### Option 2: Quick Install Script

```bash
# Linux/macOS
curl -fsSL https://raw.githubusercontent.com/m99Tanishq/CLI/main/scripts/install.sh | bash

# Or run the script directly
git clone https://github.com/m99Tanishq/CLI.git
cd CLI
./scripts/install.sh
```

### Option 3: Build from Source

```bash
git clone https://github.com/m99Tanishq/CLI.git
cd CLI

# Build and install
make install-check

# Or build manually
go build -o CLI .
sudo mv CLI /usr/local/bin/
```


### ✅ Verify Installation

After installation, verify that CLI is working:

```bash
# Check version
CLI version

# Check if CLI is available
which CLI
```

### 🗑️ Uninstallation

```bash
# If installed via the install script
./scripts/uninstall.sh

# If installed via go install
rm $(go env GOPATH)/bin/CLI

# If installed manually
sudo rm /usr/local/bin/CLI
# or
rm ~/.local/bin/CLI
```

## 🔧 Setup

1. **Set up Deployed LLM (Ollama)**:
   - Run Ollama server.
   - Set BaseURL to your endpoint.

2. **Configure CLI**:
   ```bash
   CLI config --set base_url https://your-deployed-url
   CLI config --set model phi3:mini
   CLI config --list
   ```

## 📋 Release Notes

- **Latest Release**: [v1.0.1](https://github.com/m99Tanishq/CLI/releases/latest)
- **All Releases**: [GitHub Releases](https://github.com/m99Tanishq/CLI/releases)
- **Changelog**: See release notes for detailed changes and improvements

## 🎯 Quick Start

### Chat with AI
```bash
# Start interactive chat
CLI chat

# Chat with specific model
CLI chat --model phi3:mini
```

### File Operations
```bash
# List files
CLI files list

# Read a file
CLI files read myfile.txt

# Create a file
CLI files create newfile.txt

# Write content
CLI files write newfile.txt "Hello World"
```

### Code Analysis
```bash
# Analyze code
CLI code analyze myfile.go

# Fix code issues
CLI code fix myfile.go

# Code review
CLI code review myfile.go
```

### Memory System
```bash
# Index a codebase
CLI memory index .

# Query the codebase
CLI memory query "What is the main function?"

# List indexed data
CLI memory list

# Analyze codebase
CLI memory analyze
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

The CLI stores configuration in `~/.CLI/config.json`:

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
CLI/
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
git clone https://github.com/m99Tanishq/CLI.git
cd CLI
go build -o CLI .
```

### Run Tests
```bash
go test ./...
```