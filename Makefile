.PHONY: dev build run swag test help

dev:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal
	@echo "Swagger docs generated"
	@echo "Starting dev mode (Air hot reload)..."
	air

build:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal
	@echo "Swagger docs generated"
	@echo "Building Go application..."
	@mkdir -p bin
	go build -o bin/app ./cmd/app
	@echo "Build complete: bin/app"

run:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal
	@echo "Swagger docs generated"
	@echo "Running Go application..."
	go run ./cmd/app

swag:
	@echo "Generating Swagger docs..."
	@swag init --parseDependency --parseInternal -g ./cmd/app/main.go
	@echo "Swagger docs generated"

test:
	@echo "Running tests with gotestsum..."
	gotestsum --format testname

help:
	@echo "Available commands:"
	@echo "  dev    - Start Air hot reload (with Swagger hot reload)"
	@echo "  build  - Build Go app into ./bin/app (Swagger updated first)"
	@echo "  run    - Run Go app (no reload)"
	@echo "  swag   - Generate Swagger docs"
	@echo "  test   - Run tests with gotestsum"
	@echo "  help   - Show this help"
