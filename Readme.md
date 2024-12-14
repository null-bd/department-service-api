# Go Microservice Template

A production-ready Go microservice template with built-in authentication, configuration management, and health checks.

## Features

- Clean architecture with proper separation of concerns
- Configuration management using Viper
- Authentication middleware with Keycloak integration
- PostgreSQL database integration using pgx
- Health check endpoints
- Docker and Docker Compose support
- Comprehensive test suite
- Makefile for common operations

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL
- Make (optional, for using Makefile commands)

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   ├── config.go
│   ├── database/
│   └── router/
├── internal/
│   ├── app/
│   ├── handler/
│   ├── service/
│   └── repository/
├── resources/
│   ├── config.yaml
│   ├── config-local.yaml
│   ├── config-stg.yaml
│   └── config-prod.yaml
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/your-username/your-service.git
cd your-service
```

2. Install dependencies:
```bash
go mod download
```

3. Set up configuration:
```bash
cp resources/config-local.yaml resources/config.yaml
# Edit config.yaml with your settings
```

4. Start the development environment:
```bash
make compose-up
```

5. Run the application:
```bash
make dev
```

## Configuration

The application uses a hierarchical configuration system:

1. Default configuration in `resources/config.yaml`
2. Environment-specific configurations:
   - `config-local.yaml`
   - `config-stg.yaml`
   - `config-prod.yaml`
3. Environment variables override (prefixed with `APP_`)

## Testing

Run the test suite:
```bash
make test
```

## Deployment

1. Build Docker image:
```bash
make docker-build
```

2. Run with Docker Compose:
```bash
make compose-up
```

## Available Make Commands

- `make build`: Build the application
- `make test`: Run tests
- `make run`: Run the application locally
- `make docker-build`: Build Docker image
- `make docker-run`: Run Docker container
- `make compose-up`: Start all services
- `make compose-down`: Stop all services
- `make lint`: Run linter
- `make mock`: Generate mocks for testing
- `make migrate-up`: Run database migrations
- `make migrate-down`: Rollback database migrations

## Health Check

The service provides a health check endpoint at `/health` that monitors:
- Application status
- Database connectivity
- Authentication service status

## Authentication

Authentication is handled via a auth service. Configure the following in your environment:
- Keycloak URL
- Realm
- Client ID and Secret
- Resource permissions
