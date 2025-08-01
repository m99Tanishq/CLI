#!/bin/bash

# CLI v1.0.3 Release Script
# This script helps create a GitHub release with all binaries

set -e

VERSION="v1.0.3"
RELEASE_TITLE="🚀 CLI v1.0.3 - Optimized Performance Release"
RELEASE_NOTES_FILE="RELEASE_v1.0.3.md"

echo "🎉 Creating GitHub Release for $VERSION"
echo "========================================"

# Check if we have the binaries
echo "📦 Checking for release binaries..."
BINARIES=(
    "CLI-linux-amd64"
    "CLI-linux-arm64"
    "CLI-darwin-amd64"
    "CLI-darwin-arm64"
    "CLI-windows-amd64.exe"
    "CLI-windows-arm64.exe"
)

for binary in "${BINARIES[@]}"; do
    if [ -f "$binary" ]; then
        echo "✅ Found $binary"
    else
        echo "❌ Missing $binary"
        exit 1
    fi
done

# Check if release notes exist
if [ ! -f "$RELEASE_NOTES_FILE" ]; then
    echo "❌ Release notes file $RELEASE_NOTES_FILE not found"
    exit 1
fi

echo ""
echo "📋 Release Information:"
echo "   Version: $VERSION"
echo "   Title: $RELEASE_TITLE"
echo "   Notes: $RELEASE_NOTES_FILE"
echo ""

# Read release notes
RELEASE_BODY=$(cat "$RELEASE_NOTES_FILE")

echo "🔗 GitHub Release Creation Instructions:"
echo "========================================"
echo ""
echo "1. Go to: https://github.com/m99Tanishq/CLI/releases/new"
echo ""
echo "2. Fill in the following details:"
echo "   - Tag version: $VERSION"
echo "   - Release title: $RELEASE_TITLE"
echo "   - Description: (Copy from $RELEASE_NOTES_FILE)"
echo ""
echo "3. Upload these binaries:"
for binary in "${BINARIES[@]}"; do
    echo "   - $binary"
done
echo ""
echo "4. Mark as 'Latest release'"
echo "5. Click 'Publish release'"
echo ""

echo "📝 Release Notes Preview:"
echo "========================="
echo "$RELEASE_BODY" | head -20
echo "..."
echo ""

echo "🎯 Next Steps:"
echo "=============="
echo "1. Create the release on GitHub using the instructions above"
echo "2. Upload all 6 binary files"
echo "3. Copy the release notes content"
echo "4. Publish the release"
echo ""
echo "✅ Release script completed successfully!" 