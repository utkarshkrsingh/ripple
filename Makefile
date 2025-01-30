BINARY_NAME := flowgo
BINARY_DIR := ./bin
TARGET_DIR := ./cmd/cli

deps:
	@go mod download
	@go mod tidy

build:
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(TARGET_DIR)

clean:
	@$(RM) -f $(BINARY_DIR)/$(BINARY_NAME)
