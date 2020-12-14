# Get current directory
current_dir := $(shell pwd)

# Get current commit hash
# commit_hash := $(shell git rev-parse --short=7 HEAD)

# Targets
all: testing clean build

test:
	@echo "Running all tests"

build:
	@echo "Building binaries"

	mkdir $(current_dir)/build

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build