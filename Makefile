USER := $(shell gh repo view --json owner --jq '.owner.login')
VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags="-X main.version=$(VERSION)"
# name of the binary
APP_NAME = gobpmn

# target paths
BIN_PATH = ./bin/$(APP_NAME)
INSTALL_PATH = /usr/local/bin/$(APP_NAME)

# default target
all: build

# Build for binary (./bin)
build:
	@echo "Build $(APP_NAME) with version $(VERSION)..."
	go build $(LDFLAGS) -o $(BIN_PATH) ./cmd/$(APP_NAME)

# Installation of the binary to /usr/local/bin
install: build
	@echo "Install to $(INSTALL_PATH)..."
	sudo cp $(BIN_PATH) $(INSTALL_PATH)
	@echo "Installation completed"

# Global installation via go install (requires: go install ...@version)
goinstall:
	@echo "Install globally via go install..."
	go install $(LDFLAGS) ./cmd/$(APP_NAME)
	@echo "Installed globally via go install"


update-pkg-cache:
	@echo "Update package cache..."
    GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/$(USER)/$(APP_NAME)@v$(VERSION)

# remove binary
clean:
	@echo "Clean up build directory..."
	rm -rf ./bin

.PHONY: all build install goinstall update-pkg-cache clean
