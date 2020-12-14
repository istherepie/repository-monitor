# Get current directory
current_dir := $(shell pwd)

# Namespace
namespace = github.com/istherepie/request-monitor

# Get current commit hash
commit_hash := $(shell git rev-parse --short=7 HEAD)

# Targets
.PHONY: test

all: test clean build

test: 
	@echo "Running all tests"
	go clean -testcache
	go test -v $(namespace)/webserver

build:
	@echo "Building binaries"
	mkdir $(current_dir)/build
	go build -o $(current_dir)/build/request-monitor cmd/main.go

container:
	docker build -t istherepie/request-monitor:$(commit_hash) .
	docker tag istherepie/request-monitor:$(commit_hash) istherepie/request-monitor:latest

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build