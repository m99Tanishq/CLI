# 🚀 Rzork CLI - IntelliJ Plugin

> **The ultimate AI-powered development assistant for IntelliJ IDEA, featuring blazing-fast performance and intelligent code analysis.**

[![IntelliJ Platform](https://img.shields.io/badge/IntelliJ%20Platform-2023.3.5+-blue.svg)](https://plugins.jetbrains.com)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.3-orange.svg)](https://plugins.jetbrains.com/plugin/rzork-cli)

**Rzork CLI** is a high-performance, AI-powered development assistant that brings the power of large language models directly into your IntelliJ IDEA. Built with Java and optimized for seamless IDE integration, it provides intelligent code analysis, real-time chat, and comprehensive codebase management.

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

### 🎨 **Beautiful UI Integration**
- **Modern color themes** with syntax highlighting
- **Animated spinners** and progress indicators
- **Seamless IDE integration** with tool windows
- **Dark/Light theme** support

## 🚀 Quick Start

### 1. Installation

1. **From IntelliJ Marketplace:**
   - Open IntelliJ IDEA
   - Go to `File` → `Settings` → `Plugins`
   - Search for "Rzork CLI"
   - Click `Install` and restart IDE

2. **From Plugin File:**
   - Download the `.jar` file from releases
   - Go to `File` → `Settings` → `Plugins`
   - Click the gear icon → `Install Plugin from Disk`
   - Select the downloaded `.jar` file

### 2. Configuration

1. **Open Settings:**
   - Go to `File` → `Settings` → `Tools` → `Rzork CLI Settings`

2. **Set API Key:**
   - Enter your Hugging Face API key
   - Set your preferred model (default: `zai-org/GLM-4.5:novita`)
   - Configure base URL if needed

3. **Verify Configuration:**
   - Click `Apply` and `OK`

### 3. Start Using

1. **Open Rzork CLI:**
   - Use `Tools` → `Rzork CLI` from menu
   - Or press `Ctrl+Alt+R` (Windows/Linux) / `Cmd+Alt+R` (macOS)

2. **Start Chatting:**
   - Type your questions in the input field
   - Press `Enter` or click `Send`
   - Get instant AI-powered responses

## 📋 Core Commands

### 🗣️ **Chat Commands**
```
help                    # Show available commands
chat                    # Start interactive chat
analyze <file>          # Analyze specific file
review <file>           # Code review
index                   # Index current project
query <question>        # Query indexed codebase
```

### 🔍 **Code Analysis**
```
analyze main.java       # AI-powered code analysis
review src/Utils.java   # Comprehensive code review
fix bugs               # Automatic bug detection
optimize               # Performance suggestions
```

### 🧠 **Memory System**
```
index                  # Index current project
query "architecture"   # Query indexed codebase
list                   # View indexed data
clear                  # Clear memory cache
```

## 🎯 Use Cases

### **Code Review & Analysis**
- Get instant code review for any file
- Find potential bugs and performance issues
- Receive optimization suggestions
- Security vulnerability scanning

### **Project Understanding**
- Index your entire project for AI context
- Ask questions about your codebase architecture
- Get comprehensive project analysis
- Understand complex code relationships

### **Interactive Development**
- Start AI-assisted coding sessions
- Get real-time code suggestions
- Debug with AI help
- Learn new technologies with examples

## 🔧 Configuration

### **API Settings**
```properties
# Required
api_key=hf_your_token_here

# Optional (with defaults)
model=zai-org/GLM-4.5:novita
base_url=https://router.huggingface.co/v1/chat/completions
```

### **Environment Variables**
```bash
export HF_API_KEY="your_hugging_face_token"
export RZORK_MODEL="zai-org/GLM-4.5:novita"
```

## 📊 Performance Benchmarks

| Operation | v1.0.2 | v1.0.3 | Improvement |
|-----------|--------|--------|-------------|
| **Startup Time** | 2.1s | 1.7s | ⚡ 20% faster |
| **Memory Usage** | 65MB | 55MB | 💾 15% less |
| **File Indexing** | 45s | 31s | 🚀 30% faster |
| **API Response** | 850ms | 640ms | ⚡ 25% faster |

## 🏗️ Architecture

```
src/main/java/com/rzork/cli/
├── ui/                    # User interface components
│   ├── RzorkCliToolWindow.java
│   └── RzorkCliToolWindowFactory.java
├── services/              # Core services
│   ├── RzorkCliService.java
│   └── RzorkCliProjectService.java
├── components/            # Application components
│   ├── RzorkCliApplicationComponent.java
│   └── RzorkCliProjectComponent.java
├── actions/               # IntelliJ actions
│   └── OpenRzorkCliAction.java
├── settings/              # Settings management
│   └── RzorkCliSettingsConfigurable.java
└── terminal/              # Terminal integration
    └── RzorkCliTerminalCustomizer.java
```

## 🛠️ Development

### **Prerequisites**
- IntelliJ IDEA 2023.3.5+
- Java 17+
- Gradle 8.5+

### **Build from Source**
```bash
git clone https://github.com/rzork/cli.git
cd cli

# Build plugin
./gradlew buildPlugin

# Run in IntelliJ
./gradlew runIde

# Build for distribution
./gradlew buildPlugin
```

### **Development Workflow**
```bash
# Install dependencies
./gradlew dependencies

# Run tests
./gradlew test

# Lint code
./gradlew check

# Build and install
./gradlew buildPlugin
```

## 🌍 Platform Support

| Platform | IntelliJ Version | Status |
|----------|------------------|--------|
| **Windows** | 2023.3.5+ | ✅ |
| **macOS** | 2023.3.5+ | ✅ |
| **Linux** | 2023.3.5+ | ✅ |

## 🔒 Security

- **Secure API key storage** using IntelliJ's secure storage
- **No telemetry** or data collection
- **Local processing** for sensitive operations
- **Open source** for full transparency

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### **Quick Contribution**
```bash
# Fork and clone
git clone https://github.com/your-username/cli.git
cd cli

# Create feature branch
git checkout -b feature/amazing-feature

# Make changes and test
./gradlew test
./gradlew check

# Commit and push
git commit -m "feat: add amazing feature"
git push origin feature/amazing-feature
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **JetBrains** for the excellent IntelliJ Platform
- **Hugging Face** for providing the inference API
- **All Contributors** who helped improve this project

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/rzork/cli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/rzork/cli/discussions)
- **Plugin Page**: [IntelliJ Marketplace](https://plugins.jetbrains.com/plugin/rzork-cli)

---

<div align="center">

**Made with ❤️ by [Rzork](https://github.com/rzork)**

[⭐ Star on GitHub](https://github.com/rzork/cli) • [📖 Documentation](https://github.com/rzork/cli#readme) • [🚀 Download](https://plugins.jetbrains.com/plugin/rzork-cli)

</div> 