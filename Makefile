.PHONY: all build test clean run docker-build docker-run

# Variables
APP_NAME=my-service
GO_BUILD_FLAGS=-v
DOCKER_IMAGE_NAME=my-service
DOCKER_IMAGE_TAG=latest

# Go commands
all: test build

build:
	@echo "Building application..."
	go build $(GO_BUILD_FLAGS) -o bin/$(APP_NAME) .

test:
	@echo "Running tests..."
	go test -v -race -cover ./...

clean:
	@echo "Cleaning up..."
	rm -rf bin/
	go clean -cache

run:
	@echo "Running application..."
	go run .

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

# Development commands
dev:
	@echo "Running in development mode..."
	APP_ENV=local go run ./cmd/server

lint:
	@echo "Running linter..."
	golangci-lint run

mock:
	@echo "Generating mocks..."
	mockery --all --dir internal --output internal/mocks

# Database commands
migrate-up:
	@echo "Running database migrations..."
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/myapp?sslmode=disable" up

migrate-down:
	@echo "Rolling back database migrations..."
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/myapp?sslmode=disable" down

# Docker Compose commands
compose-up:
	@echo "Starting all services..."
	docker-compose up -d

compose-down:
	@echo "Stopping all services..."
	docker-compose down