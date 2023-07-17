SHELL=/usr/bin/env bash
GO_BUILD_IMAGE?=golang:1.19

.PHONY: all
all: build

.PHONY: build
build:
	go build -o car-downloader

.PHONE: clean
clean:
	rm -f car-downloader

install:
	install -C -m 0755 car-downloader /usr/local/bin