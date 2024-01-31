GO_CMD=go
BINARY_NAME=p8r

# Mark targets as phony (not files)
.PHONY: all build clean run test

# Default target
all: build

# Build the binary
build:
	$(GO_CMD) mod tidy
	$(GO_CMD) build -ldflags="-w -s -X 'github.com/cdalar/parampiper/cmd.Version=`git rev-parse HEAD | cut -c1-7`' \
		-X 'github.com/cdalar/parampiper/cmd.BuildTime=`date -u '+%Y-%m-%d %H:%M:%S'`' \
		-X 'github.com/cdalar/parampiper/cmd.GoVersion=`go version`'" \
		-o $(BINARY_NAME) main.go

# Clean up the binary
clean:
	rm $(BINARY_NAME)

# Test the application
test:
	$(GO_CMD) test ./...
