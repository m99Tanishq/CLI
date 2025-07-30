#!/bin/bash

# Release script for CLI
# Usage: ./scripts/release.sh <version>
# Example: ./scripts/release.sh v1.0.0

set -e

if [ $# -eq 0 ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.0.0"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in format vX.Y.Z (e.g., v1.0.0)"
    exit 1
fi

echo "🚀 Creating release for version: $VERSION"

# Check if we're on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo "⚠️  Warning: You're not on the main branch (current: $CURRENT_BRANCH)"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Check for uncommitted changes
if [ -n "$(git status --porcelain)" ]; then
    echo "❌ Error: You have uncommitted changes. Please commit or stash them first."
    git status --short
    exit 1
fi

# Pull latest changes
echo "📥 Pulling latest changes..."
git pull origin main

# Run tests and security checks
echo "🧪 Running tests and security checks..."
make test-full

# Build for all platforms
echo "🔨 Building for all platforms..."
make build-all

# Create and push tag
echo "🏷️  Creating tag: $VERSION"
git tag $VERSION

echo "📤 Pushing tag to remote..."
git push origin $VERSION

echo "✅ Release process started!"
echo ""
echo "📋 Next steps:"
echo "1. Monitor the GitHub Actions workflow: https://github.com/m99Tanishq/CLI/actions"
echo "2. Wait for the release workflow to complete"
echo "3. Download the release from: https://github.com/m99Tanishq/CLI/releases"
echo ""
echo "🎉 Release $VERSION is being created automatically!" 