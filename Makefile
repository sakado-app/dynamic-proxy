# https://gist.github.com/subfuzion/0bd969d08fe0d8b5cc4b23c795854a13
# By Tony 'subfuzion' Pujals

SHELL := /bin/bash

TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: build clean fmt run

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

fmt:
	@gofmt -l -w $(SRC)

run: install
	@$(TARGET)