#!/bin/bash

# Installation script for glm-cli
# This script installs glm-cli globally and sets up PATH

set -e

echo "🚀 Installing glm-cli..."

# Check if Go is installed
if ! command -v go >/dev/null 2>&1; then
    echo "❌ Error: Go is not installed"
    echo "   Please install Go from https://golang.org/dl/"
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Get Go paths
GOPATH=$(go env GOPATH)
GOBIN=$(go env GOBIN)

# Use GOBIN if set, otherwise use GOPATH/bin
if [ -z "$GOBIN" ]; then
    INSTALL_DIR="$GOPATH/bin"
else
    INSTALL_DIR="$GOBIN"
fi

echo "📁 Installing to: $INSTALL_DIR"

# Install the CLI
echo "🔨 Building and installing glm-cli..."
go install .

# Check if installation was successful
if [ -f "$INSTALL_DIR/glm-cli" ]; then
    echo "✅ glm-cli installed successfully!"
else
    echo "❌ Installation failed"
    exit 1
fi

# Check if the binary is in PATH
if command -v glm-cli >/dev/null 2>&1; then
    echo "✅ glm-cli is available globally!"
    echo ""
    echo "🎉 Installation complete!"
    echo "   Try running: glm-cli --help"
else
    echo "⚠️  glm-cli is installed but not in PATH"
    echo ""
    echo "📝 To make it globally available, add this to your shell configuration:"
    echo ""
    
    # Detect shell and provide appropriate instructions
    SHELL_CONFIG=""
    if [ -n "$ZSH_VERSION" ]; then
        SHELL_CONFIG="$HOME/.zshrc"
        echo "   For Zsh (add to ~/.zshrc):"
    elif [ -n "$BASH_VERSION" ]; then
        SHELL_CONFIG="$HOME/.bashrc"
        echo "   For Bash (add to ~/.bashrc):"
    else
        SHELL_CONFIG="$HOME/.profile"
        echo "   For your shell (add to ~/.profile):"
    fi
    
    echo "   export PATH=\"\$PATH:$INSTALL_DIR\""
    echo ""
    echo "   Or run this command to add it temporarily:"
    echo "   export PATH=\"\$PATH:$INSTALL_DIR\""
    echo ""
    echo "   After adding to your shell config, restart your terminal or run:"
    echo "   source $SHELL_CONFIG"
fi

echo ""
echo "📚 For more information, visit: https://github.com/m99Tanishq/glm-cli" 