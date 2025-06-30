.PHONY: build

dev: serve
	@go run cmd/server/main.go

serve: 
	@go mod tidy

build:
	@go build -o build/server cmd/server/main.go
