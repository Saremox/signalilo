# Project parameters
BINARY_NAME ?= signalilo

VERSION = $(shell git describe --tags --always --dirty --match=v* || (echo "command failed $$?"; exit 1))

COMMIT=$(shell git rev-parse HEAD)
REPOSITORY ?= ghcr.io/saremox/signalilo
TAG ?= $(VERSION)

# Go parameters
GOCMD   ?= go
GOBUILD ?= $(GOCMD) build
GOCLEAN ?= $(GOCMD) clean
GOTEST  ?= $(GOCMD) test
GOGET   ?= $(GOCMD) get

.PHONY: all
all: test build

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -v \
		-o $(BINARY_NAME) \
		-ldflags "-w -s -X main.Version=$(VERSION) -X 'main.BuildDate=$(shell date -Iseconds)'"
	@echo built '$(VERSION)'

.PHONY: test
test:
	$(GOTEST) -v -cover ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf $(web_dir)

.PHONY: docker
docker:
	docker build --build-arg VERSION=$(BINARY_VERSION) -t $(REPOSITORY):$(COMMIT) .
	@echo built image $(REPOSITORY):$(COMMIT)

.PHONY: image-release
image-release:
	docker buildx build \
	--platform linux/amd64,linux/arm64,linux/arm/v7 \
	--label "org.opencontainers.image.source=https://github.com/saremox/signalilo" \
 	--label "org.opencontainers.image.description=Signalilo push alertmanager alerts to icinga2" \
 	--label "org.opencontainers.image.licenses=BSD 3-Clause" \
	--push \
	--build-arg VERSION=$(TAG) \
	-t $(REPOSITORY):latest \
	-t $(REPOSITORY):$(COMMIT) \
	-t $(REPOSITORY):$(TAG) \
	-f Dockerfile \
	.
