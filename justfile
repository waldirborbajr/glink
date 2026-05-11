#!/usr/bin/env just --justfile

set dotenv-load := true

ARTIFACT_NAME := "glink"
VERSION := `git rev-parse --short HEAD`
LDFLAGS := "-ldflags '-s -w -X main.BuildVersion={{VERSION}}'"
GCFLAGS := ""

# Default recipe - show help
default:
    @just --list

# 🔨 Build development binaries for Linux
build: tidy
    mkdir -p bin
    GOOS=linux go build -o bin/{{ARTIFACT_NAME}} {{LDFLAGS}} {{GCFLAGS}} -debug-trace=tmp/trace.json main.go

# Tidy dependencies
tidy:
    go mod tidy

# ♻️ Clean up build artifacts and caches
clean:
    rm -rf bin dist tmp
    go clean -modcache
    go clean --cache

# 🔍 Lint & format with automatic fixes
lint:
    @golangci-lint --version
    GOMEMLIMIT=1024MiB golangci-lint run -v --fix --config .golangci.yaml

# 📦 Install into GOBIN directory
install: tidy
    go install ./...

# 📝 Run tests with coverage report
test:
    go test -coverprofile cover.out -v $(go list ./... | grep -v /test/)
    go tool cover -html=cover.out

# 📸 Create snapshot release
snap: clean
    goreleaser release --snapshot

# 💻 Run Zellij with layout
layout:
    zellij --layout go-layout.kdl

# 🧑 Generate AUTHORS file from git history
authors:
    git log --format="%an" | sort | uniq > AUTHORS.txt

# 🔗 Debug build with extra symbols
debug: GCFLAGS := "-gcflags=all='-l -N'"
debug: build
