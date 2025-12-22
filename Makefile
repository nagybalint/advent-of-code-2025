.PHONY: run build

BINARY ?= aoc25
ARGS ?=

run:
	go run ./cmd ${ARGS}

build:
	go build -o $(BINARY) ./cmd