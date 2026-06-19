APP_NAME := pallet
APP_PACKAGE := ./cmd/pallet
DIST_DIR := dist
BUILD_FLAGS := -trimpath -buildvcs=false
LD_FLAGS := -s -w -buildid=

.DEFAULT_GOAL := all

.PHONY: all bench build check clean install lint test update

all: build

build:
	mkdir -p $(DIST_DIR)
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -ldflags="$(LD_FLAGS)" -o $(DIST_DIR)/$(APP_NAME) $(APP_PACKAGE)

test:
	go test -race -cover ./...

lint:
	golangci-lint run
	markdownlint-cli2 "**/*.md"
	checkmake Makefile
	yamlfmt -lint ".github/workflows/**/*.yml" ".github/workflows/**/*.yaml"

install:
	CGO_ENABLED=0 go install $(BUILD_FLAGS) -ldflags="$(LD_FLAGS)" $(APP_PACKAGE)

bench:
	go test -run="^$$" -bench=. -benchmem ./...

check: test lint

update:
	go get -u ./...
	go mod tidy

clean:
	rm -rf $(DIST_DIR)
