#!/bin/bash

set -eo pipefail

echo " --------------------------------- "
echo "|  SEVP ONE-LINER INSTALL SCRIPT  |"
echo " --------------------------------- "

# NOTIFY USER
echo
echo "[NOTICE]"
echo "This will install sevp binary to /usr/local/bin/ and add a shellhook to your shellrc."
read -p "Press [ENTER] to continue or [CTRL+C] to cancel..."

# set the repo, release tag, and binary path
BIN_NAME="sevp"
REPO="masamerc/sevp"
RELEASE_BRANCH="pre-release"
RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep "tag_name" | cut -d '"' -f 4)
RELEASE_WITHOUT_V=${RELEASE:1}
INSTALL_DIR="/usr/local/bin/"

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
    arm64)
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
echo
echo "[INFO]"
echo "OS: $os -> $TARGET_OS"
echo "Architecture: $arch -> $TARGET_ARCH"

# download and extract
echo
echo "[INSTALLATION]"
echo "Downloading and installing..."
echo "$TARBALL_URL..."
curl -sL "$TARBALL_URL" | sudo tar xz -C "$INSTALL_DIR" --wildcards "$BIN_NAME"

# check installation
if [ -x "$INSTALL_DIR$BIN_NAME" ]; then
  echo "Installation complete."
  echo "sevp installed to $INSTALL_DIR$BIN_NAME"
else
  echo "Installation failed. Please check permissions or try again."
  exit 1
fi

# install shellhook
echo
echo "[SHELLHOOK INSTALLATION]"
echo "Installing shellhook..."
curl -sSL https://raw.githubusercontent.com/masamerc/sevp/${RELEASE_BRANCH}/scripts/install_shellhook.sh | sh

echo
echo "Re-launch your shell to start using sevp!"
