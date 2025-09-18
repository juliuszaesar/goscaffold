# Go Scaffold - DDD & Clean Architecture

A production-ready Go project scaffold implementing Domain-Driven Design (DDD) and Clean Architecture principles.

[![CI/CD Pipeline](https://github.com/juliuszaesar/goscaffold/actions/workflows/ci.yml/badge.svg)](https://github.com/juliuszaesar/goscaffold/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/juliuszaesar/goscaffold)](https://goreportcard.com/report/github.com/juliuszaesar/goscaffold)
[![codecov](https://codecov.io/gh/juliuszaesar/goscaffold/branch/main/graph/badge.svg)](https://codecov.io/gh/juliuszaesar/goscaffold)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 🏗️ Architecture

This project follows Clean Architecture principles with clear separation of concerns:

```
├── cmd/api/                    # Application entry points
├── internal/
│   ├── domain/                # Domain layer (entities, value objects, interfaces)
│   │   ├── entity/           # Domain entities
│   │   ├── valueobject/      # Value objects
│   │   └── repository/       # Repository interfaces
│   ├── application/          # Application layer (use cases, DTOs)
│   │   ├── dto/              # Data Transfer Objects
│   │   └── service/          # Application services
│   ├── infrastructure/       # Infrastructure layer (external concerns)
│   │   ├── config/           # Configuration management
│   │   ├── database/         # Database connections
│   │   ├── logger/           # Logging implementation
│   │   └── repository/       # Repository implementations
│   └── interfaces/           # Interface adapters (HTTP, gRPC, etc.)
│       └── http/             # HTTP interface layer
│           ├── handler/      # HTTP handlers
│           ├── middleware/   # HTTP middleware
│           └── router/       # Route configuration
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Make (optional, for convenience commands)

### Local Development

1. **Clone the repository:**
   ```bash
   git clone https://github.com/juliuszaesar/goscaffold.git
   cd goscaffold
   ```

2. **Start with Docker Compose:**
   ```bash
   docker-compose up -d
   ```

3. **Or run locally:**
   ```bash
   # Install dependencies
   go mod download
   
   # Run the application
   make run-dev
   ```

4. **Test the API:**
   ```bash
   curl http://localhost:8080/health
   ```

## 🛠️ Development

### Available Make Commands

```bash
make help              # Show all available commands
make build             # Build the application
make test              # Run tests with coverage
make lint              # Run linter
make fmt               # Format code
make run               # Run the application
make run-dev           # Run in development mode
make docker-build      # Build Docker image
make docker-compose-up # Start all services
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage report
make coverage

# Run benchmarks
make benchmark
```

### Code Quality

The project uses several tools to maintain code quality:

- **golangci-lint**: Comprehensive linting
- **gofmt**: Code formatting
- **go vet**: Static analysis
- **staticcheck**: Advanced static analysis
- **gosec**: Security scanning

## 🔄 CI/CD Pipeline

### GitHub Actions Workflows

1. **CI/CD Pipeline** (`.github/workflows/ci.yml`):
   - Runs on push to `main`/`develop` and pull requests
   - Code quality checks (linting, formatting, security)
   - Comprehensive testing with PostgreSQL
   - Docker image building and pushing
   - Security scanning with Trivy
   - Automated deployment to staging/production

2. **Release Workflow** (`.github/workflows/release.yml`):
   - Triggered on version tags (`v*`)
   - Builds multi-platform binaries
   - Creates GitHub releases with changelogs
   - Publishes Docker images with semantic versioning

3. **Dependabot Auto-merge** (`.github/workflows/dependabot-auto-merge.yml`):
   - Automatically approves and merges minor/patch dependency updates
   - Maintains security and keeps dependencies current

### Pipeline Features

- ✅ **Multi-stage testing** with PostgreSQL integration
- ✅ **Code quality gates** (linting, formatting, security)
- ✅ **Multi-platform Docker builds** (AMD64, ARM64)
- ✅ **Security scanning** with Trivy
- ✅ **Automated dependency updates** with Dependabot
- ✅ **Semantic versioning** and automated releases
- ✅ **Coverage reporting** with Codecov
- ✅ **Environment-specific deployments**

### Creating Releases

```bash
# Create a patch release (1.0.0 -> 1.0.1)
make release-patch

# Create a minor release (1.0.0 -> 1.1.0)
make release-minor

# Create a major release (1.0.0 -> 2.0.0)
make release-major
```

## 🐳 Docker

### Building the Image

```bash
# Build locally
docker build -t goscaffold:latest .

# Or use make
make docker-build
```

### Running with Docker

```bash
# Run single container
docker run -p 8080:8080 goscaffold:latest

# Run with full stack
docker-compose up -d
```

### Multi-stage Build

The Dockerfile uses a multi-stage build for optimal image size:
- **Builder stage**: Compiles the Go binary
- **Final stage**: Minimal scratch image with just the binary
- **Security**: Runs as non-root user
- **Health checks**: Built-in health check endpoint

## 📊 Monitoring

The docker-compose setup includes:

- **Prometheus**: Metrics collection (`:9090`)
- **Grafana**: Visualization dashboard (`:3000`)
- **PostgreSQL**: Primary database (`:5432`)
- **Redis**: Caching layer (`:6379`)

## 🔧 Configuration

Configuration is managed through environment variables:

```bash
# Server configuration
SERVER_PORT=8080
SERVER_READ_TIMEOUT=30
SERVER_WRITE_TIMEOUT=30
SERVER_IDLE_TIMEOUT=120

# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=goscaffold
DB_PASSWORD=password
DB_NAME=goscaffold
DB_SSL_MODE=disable

# Application configuration
ENVIRONMENT=development
LOG_LEVEL=info
```

## 🧪 Testing Strategy

- **Unit Tests**: Domain logic and business rules
- **Integration Tests**: Database and external service interactions
- **API Tests**: HTTP endpoint testing
- **Security Tests**: Vulnerability scanning
- **Performance Tests**: Benchmarking critical paths

## 📈 Performance

- **Graceful shutdown**: Proper cleanup on termination
- **Connection pooling**: Efficient database connections
- **Middleware stack**: Logging, CORS, recovery
- **Health checks**: Kubernetes-ready health endpoints
- **Metrics**: Prometheus-compatible metrics

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes following the coding standards
4. Run tests: `make test`
5. Run linting: `make lint`
6. Commit your changes: `git commit -m 'Add amazing feature'`
7. Push to the branch: `git push origin feature/amazing-feature`
8. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Clean Architecture by Robert C. Martin
- Domain-Driven Design by Eric Evans
- Go community best practices and idioms
