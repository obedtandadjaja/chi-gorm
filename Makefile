PROJECT := $(shell basename "$(PWD)")
BRANCH := $(shell git rev-parse --abbref-ref HEAD)
BUILD := $(shell git rev-parse --short HEAD)

run:
	@echo "  > $(PROJECT) is running..."
	@go run main.go
