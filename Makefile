ARTIFACT_NAME:=glink

GCFLAGS :=
debug: GCFLAGS += -gcflags=all='-l -N'

VERSION ?= $(shell git rev-parse --short HEAD)
LDFLAGS = -ldflags '-s -w -X main.version=$(VERSION)'

help: ## 💬 This help message :)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## 🔨 Build development binaries for Linux
	@go mod tidy
	GOOS=linux go build -o bin/$(ARTIFACT_NAME) $(LDFLAGS) $(GCFLAGS) -debug-trace=tmp/trace.json main.go

clean: ## ♻️  Clean up
	@rm -rf bin
	@rm $(GOBIN)/$(ARTIFACT_NAME)

lint-fix: ## 🔍 Lint & format, will try to fix errors and modify code
	golangci-lint --version
	GOMEMLIMIT=1024MiB @golangci-lint run -v --modules-download-mode=mod *.go --fix

install: ## Install into GOBIN directory
	@go install ./...

test: ## 📝 Run all tests
	@go test -coverprofile cover.out -v $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out

snap:
	@rm -rf dist/
	@goreleaser release --snapshot

layout: ## 💻 Run Zellij with a layout
	@zellij --layout go-layout.kdl

.PHONY: authors
authors:
	git log --format="%an" | sort | uniq > AUTHORS.txt
