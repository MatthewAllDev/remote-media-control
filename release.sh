#!/bin/bash

create_o_file() {
  VERSION_NUMBER="${VERSION//v/}"
  RC_VERSION="${VERSION_NUMBER//./,}"
  VERSION_PARTS=(${VERSION_NUMBER//./ })
  FILEVERSION="${VERSION_PARTS[0]},${VERSION_PARTS[1]},0,0"
  RC_FILE="resource.rc"
  O_FILE="version.syso"
  cat > $RC_FILE <<EOL
MAINICON ICON "icons/app.ico"

1 VERSIONINFO
FILEVERSION     ${FILEVERSION}
PRODUCTVERSION  ${FILEVERSION}
FILEOS          0x40004L
FILETYPE        0x1L
{
  BLOCK "StringFileInfo"
  {
    BLOCK "040904E4"
    {
      VALUE "CompanyName",      "Ilia Kuvarzin\0"
      VALUE "FileDescription",   "Remote Media Control\0"
      VALUE "FileVersion",       "${VERSION_NUMBER}.0\0"
      VALUE "LegalCopyright",    "Copyright © 2025 Ilia Kuvarzin. Licensed under MIT License\0"
      VALUE "ProductName",       "Remote Media Control\0"
      VALUE "ProductVersion",    "${VERSION_NUMBER}\0"
    }
  }
  BLOCK "VarFileInfo"
  {
    VALUE "Translation", 0x409, 0x4E4
  }
}
EOL
  echo ".RC file created at $RC_FILE"
  x86_64-w64-mingw32-windres $RC_FILE -O coff -o $O_FILE
  echo ".O file created at $O_FILE"
}

compile() {
  echo "Compiling for Linux..."
  GOOS=linux GOARCH=amd64 go build -o "$LINUX_DIR/$APP_NAME" -ldflags "-s -w" main.go

  echo "Compiling for Windows..."
  GOOS=windows GOARCH=amd64 go build -o "$WINDOWS_DIR/$APP_NAME.exe" -ldflags="-s -w"
}

copy_files() {
  echo "Copying setup scripts and static assets..."
  cp scripts/setup.sh "$LINUX_DIR/"
  cp scripts/start.sh "$LINUX_DIR/"
  cp scripts/uninstall.bat "$WINDOWS_DIR/"
  cp scripts/install.bat "$WINDOWS_DIR/"
  cp icons/app-min.svg "$LINUX_DIR/app.svg"
  cp -r static "$LINUX_DIR/"
  cp -r static "$WINDOWS_DIR/"
}

create_archives() {
  echo "Creating tar.gz for Linux..."
  tar -czf "$BUILD_DIR/${APP_NAME}_linux_${VERSION}.tar.gz" -C "$LINUX_DIR" .

  echo "Creating zip for Windows..."
  cd "$WINDOWS_DIR"
  zip -r "../${APP_NAME}_windows_${VERSION}.zip" .
  cd - > /dev/null
}

set -e

APP_NAME="RemoteMediaControl"
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.1")

BUILD_DIR="release"
LINUX_DIR="$BUILD_DIR/${APP_NAME}_linux"
WINDOWS_DIR="$BUILD_DIR/${APP_NAME}_windows"

echo "Building release $VERSION..."

rm -rf "$BUILD_DIR"
mkdir -p "$LINUX_DIR" "$WINDOWS_DIR"

create_o_file
compile
copy_files
create_archives

echo "Cleaning up..."
rm -rf "$LINUX_DIR" "$WINDOWS_DIR"

echo "Release build completed:"
ls -lh "$BUILD_DIR"
