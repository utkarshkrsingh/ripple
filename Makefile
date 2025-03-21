BINARY_NAME := ripple
BINARY_DIR := ./bin
TARGET_DIR := ./cmd/ripple

deps:
	@go mod tidy
	@go mod download

build:
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(TARGET_DIR)

clean:
	@$(RM) -f $(BINARY_DIR)/$(BINARY_NAME)
