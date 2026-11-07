TOOL = go tool -modfile=tools/go.mod

.PHONY: format lint test build run

format:
	$(TOOL) golangci-lint fmt

lint:
	go mod tidy -diff
	cd tools && go mod tidy -diff
	$(TOOL) golangci-lint run
	$(TOOL) govulncheck ./...

test:
	go test -race -coverprofile=cover.out ./...
	go tool cover -func=cover.out | awk '$$1 ~ /\/main\.go:/ && $$2 == "main" { next } $$1 != "total:" && $$3 != "100.0%" { print; low = 1 } END { exit low }'

build:
	docker build -t app .

run:
	go run ./cmd/app
