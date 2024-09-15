APP_NAME = ollama-dbus
GO_FILES = $(shell find . -type f -name '*.go')
GO_MOD_FILE = go.mod
BUILD_DIR = ./bin
BUILD_OUTPUT = $(BUILD_DIR)/$(APP_NAME)

all: build

build: $(BUILD_OUTPUT)

$(BUILD_OUTPUT): $(GO_FILES) $(GO_MOD_FILE)
	@echo "Building $(APP_NAME)..."
	@mkdir $(BUILD_DIR)
	@go build -o $(BUILD_OUTPUT) ./cmd/ollama-dbus

run: $(BUILD_OUTPUT)
	@echo "Running $(APP_NAME)..."
	@$(BUILD_OUTPUT)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

deps:
	@echo "Installing dependencies..."
	@go mod tidy

help:
	@echo "Makefile for Ollama D-Bus"
	@echo "Usage:"
	@echo "  make          - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make clean    - Clean up build artifacts"
	@echo "  make deps     - Install dependencies"
	@echo "  make help     - Display this help message"

.PHONY: all build run clean deps help
