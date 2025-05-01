#!/bin/bash

APP_NAME="gobpmn"
BIN_PATH="./bin/$APP_NAME"
INSTALL_PATH="/usr/local/bin/$APP_NAME"

echo "Build $APP_NAME..."
go build -o "$BIN_PATH" ./cmd/$APP_NAME

if [ $? -ne 0 ]; then
  echo "Build failed!"
  exit 1
fi

echo "Install to $INSTALL_PATH..."
sudo cp "$BIN_PATH" "$INSTALL_PATH"

if [ $? -eq 0 ]; then
  echo "$APP_NAME successfully installed!"
else
  echo "Installation failed!"
fi
