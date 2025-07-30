#!/bin/bash

# Uninstall script for CLI

set -e

echo "🗑️  Uninstalling CLI..."

# Get Go paths
GOPATH=$(go env GOPATH)
GOBIN=$(go env GOBIN)

# Use GOBIN if set, otherwise use GOPATH/bin
if [ -z "$GOBIN" ]; then
    INSTALL_DIR="$GOPATH/bin"
else
    INSTALL_DIR="$GOBIN"
fi

BINARY_PATH="$INSTALL_DIR/CLI"

echo "📁 Looking for CLI in: $INSTALL_DIR"

# Check if the binary exists
if [ -f "$BINARY_PATH" ]; then
    echo "✅ Found CLI at: $BINARY_PATH"
    
    # Remove the binary
    rm "$BINARY_PATH"
    echo "✅ Removed CLI binary"
    
    # Check if there are any other Go binaries in the directory
    if [ -z "$(ls -A "$INSTALL_DIR" 2>/dev/null)" ]; then
        echo "📁 Directory is empty, removing: $INSTALL_DIR"
        rmdir "$INSTALL_DIR"
    fi
else
    echo "❌ CLI not found at: $BINARY_PATH"
    
    # Try to find it elsewhere
    if command -v CLI >/dev/null 2>&1; then
        FOUND_PATH=$(which CLI)
        echo "📍 Found CLI at: $FOUND_PATH"
        echo "   Please remove it manually:"
        echo "   sudo rm $FOUND_PATH"
    else
        echo "✅ CLI is not installed or not in PATH"
    fi
fi

echo ""
echo "🧹 Cleaning up configuration files..."
CONFIG_DIR="$HOME/.CLI"

if [ -d "$CONFIG_DIR" ]; then
    echo "📁 Found config directory: $CONFIG_DIR"
    read -p "   Remove configuration files? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -rf "$CONFIG_DIR"
        echo "✅ Removed configuration files"
    else
        echo "📁 Configuration files preserved at: $CONFIG_DIR"
    fi
else
    echo "✅ No configuration files found"
fi

echo ""
echo "🎉 Uninstallation complete!"
echo "   If you added PATH modifications to your shell config,"
echo "   you may want to remove them manually." 