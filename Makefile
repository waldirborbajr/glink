build-dev:
	go build -o bin/glink -debug-trace=tmp/trace.json main.go

build:
	go build -o bin/glink -ldflags="-s -w" main.go

install:
	go install ./...

test:
	go test ./...

layout:
	zellij --layout go-layout.kdl
