.PHONY: build run

GIT_TAG=$(shell git describe --tags --abbrev=0 2>/dev/null || echo dev)

build:
	@echo "Building Docker image with GIT_TAG=$(GIT_TAG)"
	@GIT_TAG=$(GIT_TAG) docker compose build
run: build
	@GIT_TAG=$(GIT_TAG) docker compose up -d
