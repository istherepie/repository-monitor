# Get current directory
current_dir := $(shell pwd)

# Namespace
namespace = github.com/istherepie/request-monitor

# Get current commit hash
# commit_hash := $(shell git rev-parse --short=7 HEAD)

# Targets
all: testing clean build

test:
	@echo "Running all tests"
	go clean -testcache
	go test -v $(namespace)/webserver

build:
	@echo "Building binaries"

	mkdir $(current_dir)/build

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build