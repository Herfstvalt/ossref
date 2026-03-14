#!/bin/sh
set -e

REPO="herfstvalt/ossref"
INSTALL_DIR="/usr/local/bin"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac

case "$OS" in
  darwin|linux) ;;
  *) echo "Unsupported OS: $OS" && exit 1 ;;
esac

# Get latest version
VERSION=$(curl -sI "https://github.com/$REPO/releases/latest" | grep -i "^location:" | sed 's/.*tag\///' | tr -d '\r\n')

if [ -z "$VERSION" ]; then
  echo "Failed to detect latest version"
  exit 1
fi

URL="https://github.com/$REPO/releases/download/$VERSION/ossref_${OS}_${ARCH}.tar.gz"

echo "Installing ossref $VERSION ($OS/$ARCH)..."

TMPDIR=$(mktemp -d)
curl -sL "$URL" -o "$TMPDIR/ossref.tar.gz"
tar -xzf "$TMPDIR/ossref.tar.gz" -C "$TMPDIR"

if [ -w "$INSTALL_DIR" ]; then
  mv "$TMPDIR/ossref" "$INSTALL_DIR/ossref"
else
  sudo mv "$TMPDIR/ossref" "$INSTALL_DIR/ossref"
fi

rm -rf "$TMPDIR"

echo "✓ ossref installed to $INSTALL_DIR/ossref"
echo ""
ossref help
