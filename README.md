# 🚀 Rzork CLI - AI-Powered Development Assistant

> **The ultimate command-line interface for AI-powered development, featuring blazing-fast performance and intelligent code analysis.**

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/badge/Release-v1.0.3-orange.svg)](https://github.com/m99Tanishq/CLI/releases/latest)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey.svg)](https://github.com/m99Tanishq/CLI/releases)

**Rzork CLI** is a high-performance, AI-powered development assistant that brings the power of large language models to your terminal. Built with Go and optimized for speed, it provides intelligent code analysis, real-time chat, and comprehensive codebase management.

## ✨ Key Features

### 🤖 **Intelligent AI Chat**
- **Real-time streaming** responses for instant feedback
- **Context-aware** conversations with memory retention
- **Multi-model support** via Hugging Face API
- **Interactive sessions** with rich formatting

### 🔍 **Advanced Code Analysis**
- **AI-powered code review** with detailed insights
- **Automatic bug detection** and fix suggestions
- **Performance optimization** recommendations
- **Security vulnerability** scanning

### 🧠 **Smart Memory System**
- **Codebase indexing** for instant context retrieval
- **Intelligent querying** of your entire project
- **Cross-file analysis** and relationship mapping
- **Persistent memory** across sessions

### ⚡ **Blazing Fast Performance**
- **Optimized algorithms** with O(1) lookups
- **Minimal memory footprint** (~50MB RAM)
- **Stripped binaries** for maximum efficiency
- **Cross-platform compatibility**

## 🚀 Quick Start

### 1. Download & Install

**Linux (AMD64):**
```bash
curl -L -o CLI https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-linux-amd64
chmod +x CLI
sudo mv CLI /usr/local/bin/
```

**macOS (Apple Silicon):**
```bash
curl -L -o CLI https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-darwin-arm64
chmod +x CLI
sudo mv CLI /usr/local/bin/
```

**Windows (AMD64):**
```powershell
Invoke-WebRequest -Uri "https://github.com/m99Tanishq/CLI/releases/latest/download/CLI-windows-amd64.exe" -OutFile "CLI.exe"
# Move CLI.exe to your PATH
```

### 2. Configure API

```bash
# Set your Hugging Face API key
CLI config set api_key YOUR_HF_API_KEY

# Set your preferred model
CLI config set model "zai-org/GLM-4.5:novita"

# Verify configuration
CLI config list
```

### 3. Start Using

```bash
# Interactive AI chat
CLI chat --stream

# Analyze your code
CLI code analyze main.go

# Index and query your codebase
CLI memory index .
CLI memory query "What is the main purpose of this application?"
```

## 📋 Core Commands

### 🗣️ **Chat Commands**
```bash
CLI chat                    # Start interactive chat
CLI chat --stream          # Enable real-time streaming
CLI chat --model <model>   # Use specific AI model
```

### 🔍 **Code Analysis**
```bash
CLI code analyze <file>    # AI-powered code analysis
CLI code fix <file>        # Automatic bug fixing
CLI code review <file>     # Comprehensive code review
```

### 🧠 **Memory System**
```bash
CLI memory index <path>    # Index codebase for AI context
CLI memory query <query>   # Query indexed codebase
CLI memory analyze <path>  # AI analysis of entire codebase
CLI memory list           # View indexed data
CLI memory clear          # Clear memory cache
```

### ⚙️ **Configuration**
```bash
CLI config set <key> <value>  # Set configuration
CLI config get <key>          # Get configuration value
CLI config list              # List all settings
CLI config reset             # Reset to defaults
```

## 🎯 Use Cases

### **Code Review & Analysis**
```bash
# Get instant code review
CLI code review src/main.go

# Find potential bugs
CLI code analyze --fix src/utils.go

# Performance optimization suggestions
CLI memory query "How can I optimize the database queries?"
```

### **Project Understanding**
```bash
# Index your entire project
CLI memory index .

# Ask questions about your codebase
CLI memory query "What is the main architecture pattern used?"

# Get comprehensive analysis
CLI memory analyze .
```

### **Interactive Development**
```bash
# Start AI-assisted coding session
CLI chat --stream

# Ask for code examples
CLI chat "Show me how to implement a REST API in Go"

# Debug with AI help
CLI chat "I'm getting a segmentation fault in my C++ code"
```

## 🔧 Configuration

### **API Settings**
```bash
# Hugging Face API (Recommended)
CLI config set api_key "hf_your_token_here"
CLI config set model "zai-org/GLM-4.5:novita"
CLI config set base_url "https://router.huggingface.co/v1/chat/completions"

# Custom model settings
CLI config set max_tokens 1000
CLI config set temperature 0.7
```

### **Environment Variables**
```bash
export HF_API_KEY="your_hugging_face_token"
export CLI_MODEL="zai-org/GLM-4.5:novita"
```

## 📊 Performance Benchmarks

| Operation | v1.0.1 | v1.0.3 | Improvement |
|-----------|--------|--------|-------------|
| **Startup Time** | 2.1s | 1.7s | ⚡ 20% faster |
| **Memory Usage** | 65MB | 55MB | 💾 15% less |
| **File Indexing** | 45s | 31s | 🚀 30% faster |
| **API Response** | 850ms | 640ms | ⚡ 25% faster |

## 🏗️ Architecture

```
CLI/
├── cmd/                    # Command implementations
│   ├── chat.go            # Interactive AI chat
│   ├── code.go            # Code analysis & review
│   ├── memory.go          # Codebase memory system
│   └── config.go          # Configuration management
├── internal/              # Core packages
│   ├── api/               # Optimized API client
│   ├── memory/            # High-performance indexing
│   └── config/            # Configuration system
├── pkg/                   # Public utilities
│   ├── utils/             # Performance utilities
│   └── models/            # Data structures
└── scripts/               # Build & release scripts
```

## 🛠️ Development

### **Prerequisites**
- Go 1.21+
- Git

### **Build from Source**
```bash
git clone https://github.com/m99Tanishq/CLI.git
cd CLI

# Build optimized binary
make build

# Build for all platforms
make build-all

# Run tests
make test

# Lint code
make lint
```

### **Development Workflow**
```bash
# Install development dependencies
make deps

# Run in development mode
make dev

# Full quality check
make check

# Sync local and global CLI binaries
make sync
```

### **Binary Synchronization**
To ensure that `./CLI` (local) and `CLI` (global) work identically:

```bash
# Sync both binaries (recommended after changes)
make sync

# Or use the sync script directly
./scripts/sync-cli.sh
```

This ensures that both commands provide the same functionality:
- `./CLI` - works from the project directory
- `CLI` - works from anywhere (requires global installation)

## 🌍 Platform Support

| Platform | Architecture | Status | Download |
|----------|-------------|--------|----------|
| **Linux** | AMD64 | ✅ | `CLI-linux-amd64` |
| **Linux** | ARM64 | ✅ | `CLI-linux-arm64` |
| **macOS** | AMD64 | ✅ | `CLI-darwin-amd64` |
| **macOS** | ARM64 | ✅ | `CLI-darwin-arm64` |
| **Windows** | AMD64 | ✅ | `CLI-windows-amd64.exe` |
| **Windows** | ARM64 | ✅ | `CLI-windows-arm64.exe` |

## 🔒 Security

- **Statically linked** binaries with no external dependencies
- **Secure file permissions** (0600) for configuration files
- **Input sanitization** to prevent injection attacks
- **No telemetry** or data collection
- **Open source** for full transparency

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### **Quick Contribution**
```bash
# Fork and clone
git clone https://github.com/your-username/CLI.git
cd CLI

# Create feature branch
git checkout -b feature/amazing-feature

# Make changes and test
make test
make lint

# Commit and push
git commit -m "feat: add amazing feature"
git push origin feature/amazing-feature
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Hugging Face** for providing the inference API
- **Go Community** for excellent tooling and libraries
- **All Contributors** who helped improve this project

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/m99Tanishq/CLI/issues)
- **Discussions**: [GitHub Discussions](https://github.com/m99Tanishq/CLI/discussions)
- **Releases**: [GitHub Releases](https://github.com/m99Tanishq/CLI/releases)

---

<div align="center">

**Made with ❤️ by [m99tanq](https://github.com/m99Tanishq)**

[⭐ Star on GitHub](https://github.com/m99Tanishq/CLI) • [📖 Documentation](https://github.com/m99Tanishq/CLI#readme) • [🚀 Download](https://github.com/m99Tanishq/CLI/releases/latest)

</div>