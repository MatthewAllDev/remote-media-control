#!/bin/bash

set -e

APP_NAME="RemoteMediaControl"
VERSION=$(git describe --tags --abbrev=0)

BUILD_DIR="release"
LINUX_DIR="$BUILD_DIR/${APP_NAME}_linux"
WINDOWS_DIR="$BUILD_DIR/${APP_NAME}_windows"

echo "Building release $VERSION..."

rm -rf "$BUILD_DIR"
mkdir -p "$LINUX_DIR" "$WINDOWS_DIR"

echo "Compiling for Linux..."
GOOS=linux GOARCH=amd64 go build -o "$LINUX_DIR/$APP_NAME" main.go

echo "Compiling for Windows..."
GOOS=windows GOARCH=amd64 go build -o "$WINDOWS_DIR/$APP_NAME.exe" main.go

echo "Copying setup scripts and static assets..."
cp setup_linux.sh "$LINUX_DIR/"
cp setup_windows.ps1 "$WINDOWS_DIR/"
cp -r static "$LINUX_DIR/"
cp -r static "$WINDOWS_DIR/"

echo "Creating tar.gz for Linux..."
tar -czf "$BUILD_DIR/${APP_NAME}_linux_${VERSION}.tar.gz" -C "$LINUX_DIR" .

echo "Creating zip for Windows..."
cd "$WINDOWS_DIR"
zip -r "../${APP_NAME}_windows_${VERSION}.zip" .
cd - > /dev/null

echo "Cleaning up..."
rm -rf "$LINUX_DIR" "$WINDOWS_DIR"

echo "Release build completed:"
ls -lh "$BUILD_DIR"
