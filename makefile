BINARY_NAME=todo-api
BIN_DIR=bin
CMD_DIR=cmd
MAIN_FILE=main.go

all: build

build:
	@echo "> Building the binary..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_DIR)/$(MAIN_FILE)
	@echo "> Build complete. Binary is located at $(BIN_DIR)/$(BINARY_NAME)"

run: build
	@echo "> Running the binary..."
	@./$(BIN_DIR)/$(BINARY_NAME)

clean:
	@echo "> Cleaning up..."
	@rm -rf $(BIN_DIR)
	@echo "> Clean complete."

test:
	@echo "> Running tests..."
	@go test ./...

lint:
	@echo "> Linting the code..."
	@golangci-lint run

.PHONY: all build run clean test lint
