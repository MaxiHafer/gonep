

.PHONY: test
test:
	go test -race ./...

.PHONY: dep
dep:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2