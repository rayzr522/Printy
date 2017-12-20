#!/bin/bash
cd "$(dirname "$0")/.."

source settings.cfg

ARCH="$(uname | tr "[:upper:]" "[:lower:]")"
FILE_NAME="$NAME-$ARCH-$VERSION"
OUTPUT="dist/$FILE_NAME"

set -e

echo -n "Cleaning up old builds... "
rm -rf ./dist
echo "done"

echo -n "Building main.go... "
go build -o "$OUTPUT" main.go
echo "done"

echo -n "Linking to $BIN_LINK... "
ln -i "$OUTPUT" "$BIN_LINK"
echo "done"
