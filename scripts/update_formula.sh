#!/bin/bash

set -e

REPO="masamerc/sevp"
FORMULA_PATH="HomebrewFormula/sevp.rb"
LATEST_VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | jq -r .tag_name | sed 's/v//')
PLATFORMS=(
  "darwin_arm64"
  "darwin_amd64"
  "linux_amd64"
  "linux_arm64"
)

# update the formula file
echo "Updating formula to version $LATEST_VERSION..."

for PLATFORM in "${PLATFORMS[@]}"; do
  FILE_NAME="sevp_${LATEST_VERSION}_${PLATFORM}.tar.gz"
  URL="https://github.com/$REPO/releases/download/v$LATEST_VERSION/$FILE_NAME"

  echo "Fetching SHA256 for $PLATFORM..."
  SHA256=$(curl -sL "$URL" | sha256sum | awk '{ print $1 }')

  echo "Updating $PLATFORM in formula..."
  sed -i.bak -E "s#(url \")[^\"]*${PLATFORM}.tar.gz\"#\1$URL\"#" "$FORMULA_PATH"
  sed -i.bak -E "/${PLATFORM}.tar.gz\"/{n;s#(sha256 \")[^\"]*\"#\1$SHA256\"#}" "$FORMULA_PATH"
done

rm -f "$FORMULA_PATH.bak"

echo "Formula updated successfully"
echo "New version: $LATEST_VERSION"
