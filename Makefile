.PHONY: build
build:
	go build -o build/gonep-cli ./cmd/gonep-cli/gonep.go

.PHONY: test
test:
	go test -race ./...