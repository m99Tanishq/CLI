#!/bin/bash

# Rzork CLI IntelliJ Plugin Publish Script
# This script publishes the plugin to the IntelliJ Marketplace

set -e

echo "🚀 Publishing Rzork CLI IntelliJ Plugin..."

# Check if we're in the right directory
if [ ! -f "build.gradle.kts" ]; then
    echo "❌ Error: build.gradle.kts not found. Please run this script from the project root."
    exit 1
fi

# Check if PUBLISH_TOKEN is set
if [ -z "$PUBLISH_TOKEN" ]; then
    echo "❌ Error: PUBLISH_TOKEN environment variable is not set."
    echo "Please set your IntelliJ Marketplace publish token:"
    echo "export PUBLISH_TOKEN=your_token_here"
    exit 1
fi

# Check if certificate chain is set (for signing)
if [ -z "$CERTIFICATE_CHAIN" ]; then
    echo "⚠️  Warning: CERTIFICATE_CHAIN not set. Plugin will not be signed."
    echo "To sign the plugin, set:"
    echo "export CERTIFICATE_CHAIN=your_certificate_chain"
    echo "export PRIVATE_KEY=your_private_key"
    echo "export PRIVATE_KEY_PASSWORD=your_password"
fi

# Make gradlew executable
chmod +x gradlew

# Build the plugin first
echo "🔨 Building plugin..."
./gradlew buildPlugin

# Sign the plugin if certificate is available
if [ -n "$CERTIFICATE_CHAIN" ] && [ -n "$PRIVATE_KEY" ] && [ -n "$PRIVATE_KEY_PASSWORD" ]; then
    echo "🔐 Signing plugin..."
    ./gradlew signPlugin
else
    echo "⚠️  Skipping plugin signing (certificate not available)"
fi

# Publish to IntelliJ Marketplace
echo "📤 Publishing to IntelliJ Marketplace..."
./gradlew publishPlugin

if [ $? -eq 0 ]; then
    echo "✅ Plugin published successfully!"
    echo "🌐 Check your plugin at: https://plugins.jetbrains.com/plugin/rzork-cli"
else
    echo "❌ Publishing failed!"
    exit 1
fi

echo "�� Publish complete!" 