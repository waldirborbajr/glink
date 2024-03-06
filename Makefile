help: ## 💬 This help message :)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build-dev: ## 🔨 Build development binaries for Linux
	go mod tidy
	GOOS=linux go build -o bin/glink -debug-trace=tmp/trace.json main.go

build: ## 🔨 Build binaries for Linux
	go mod tidy
	GOOS=linux go build -o bin/glink -ldflags="-s -w" main.go

clean: ## ♻️  Clean up
	@rm -rf bin

lint-fix: ## 🔍 Lint & format, will try to fix errors and modify code
	golangci-lint run --modules-download-mode=mod *.go --fix

install: ## Install into GOBIN directory
	go install ./...

test: ## 📝 Run all tests
	go test ./...

snap:
	@rm -rf dist/
	goreleaser release --snapshot

layout: ## 💻 Run Zellij with a layout
	zellij --layout go-layout.kdl
