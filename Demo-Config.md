```toml
[variables]
BINARY_NAME = "ripple"
BINARY_DIR = "./bin"
TARGET_DIR = "./cmd/ripple"

[deps]
desc = "Package checking"
cmd = "go mod download && go mod tidy"

[build]
desc = "Building the application"
cmd = "go build -o ${BINARY_DIR}/${BINARY_NAME} ${TARGET_DIR}"
depends_on = [
    "deps"
]

[clean]
desc = "Cleaning the binary"
cmd = "rm -rf ${BINARY_DIR}/${BINARY_NAME}"
```
