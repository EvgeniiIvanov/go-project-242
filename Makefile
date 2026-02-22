.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix 
lint-fix: lint
	golangci-lint run --fix

.PHONY: test
test:
	go test ./... -v

.PHONY: build
build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
