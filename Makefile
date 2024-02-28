build-dev:
	go mod tidy
	GOOS=linux go build -o bin/glink -debug-trace=tmp/trace.json main.go

build:
	go mod tidy
	GOOS=linux go build -o bin/glink -ldflags="-s -w" main.go

lint-fix: ## 🔍 Lint & format, will try to fix errors and modify code
	golangci-lint run --modules-download-mode=mod *.go --fix

install:
	go install ./...

test:
	go test ./...

layout:
	zellij --layout go-layout.kdl
