.PHONY: build install docker-image

TENDERMINT_VERSION?=latest
BUILD_TAGS?='linkis'
VERSION := $(shell git describe --always)
BUILD_FLAGS = -ldflags "-X github.com/DeAI-Artist/Linkis/version.TMCoreSemVer=$(VERSION)

.DEFAULT_GOAL := build

build:
	CGO_ENABLED=1 go build $(BUILD_FLAGS) -tags $(BUILD_TAGS) -o ../../build/tm-signer-harness main.go

install:
	CGO_ENABLED=1 go install $(BUILD_FLAGS) -tags $(BUILD_TAGS) .

docker-image:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -tags $(BUILD_TAGS) -o tm-signer-harness main.go
	docker build \
		--build-arg TENDERMINT_VERSION=$(TENDERMINT_VERSION) \
		-t tendermint/tm-signer-harness:$(TENDERMINT_VERSION) .
	rm -rf tm-signer-harness
