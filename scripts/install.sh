#!/bin/bash

set -eo pipefail

echo " --------------------------------- "
echo "|  SEVP ONE-LINER INSTALL SCRIPT  |"
echo " --------------------------------- "

# set the repo, release tag, and binary path
BIN_NAME="sevp"
REPO="masamerc/sevp"
RELEASE_BRANCH="pre-release"
RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep "tag_name" | cut -d '"' -f 4)
RELEASE_WITHOUT_V=${RELEASE:1}
INSTALL_DIR="/usr/local/bin"

# os
os=$(uname -s)
case "$os" in
    Linux)
        TARGET_OS="linux"
    ;;
    Darwin)
        TARGET_OS="darwin"
    ;;
    *)
        echo "Unsupported os for this one-liner install script: $arch"
        exit 1
    ;;
esac

# architecture
arch=$(uname -m)
case "$arch" in
    x86_64)
        TARGET_ARCH="amd64"
    ;;
    aarch64)
        TARGET_ARCH="arm64"
    ;;
    *)
        echo "Unsupported architecture for this one-liner install script: $arch"
        exit 1
    ;;
esac

# construct the download URL
TARBALL_URL="https://github.com/$REPO/releases/download/$RELEASE/sevp_${RELEASE_WITHOUT_V}_${TARGET_OS}_${TARGET_ARCH}.tar.gz"

# start installation
echo "OS: $os -> $TARGET_OS"
echo "Architecture: $arch -> $TARGET_ARCH"

# download and extract
echo
echo "Downloading..."
echo "$TARBALL_URL..."
echo "Download completed."

curl -sL "$TARBALL_URL" | tar xz

# move binary to the install directory (adjust the binary name as needed)
echo
echo "Installing..."
sudo mv $BIN_NAME "$INSTALL_DIR/$BIN_NAME"
echo "Installed to $INSTALL_DIR/$BIN_NAME"


# check installation
if [ -x "$INSTALL_DIR/$BIN_NAME" ]; then
  echo "Installation complete."
else
  echo "Installation failed. Please check permissions or try again."
  exit 1
fi

# install shellhook
echo
echo "Installing shellhook..."
curl -sSL https://raw.githubusercontent.com/masamerc/sevp/${RELEASE_BRANCH}/scripts/install_shellhook.sh | sh
echo "Shellhook installed for your current shell."
