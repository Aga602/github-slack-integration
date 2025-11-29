# GitHub Slack Integration

A Go application for integrating GitHub with Slack notifications.

## Project Structure

```
.
├── api/                    # API definitions and OpenAPI specs
├── cmd/
│   └── app/                # Application entry point and setup
├── configs/                # Configuration files and examples
├── internal/               # Private application code
│   ├── config/             # Configuration loading
│   ├── handler/            # HTTP handlers
│   ├── model/              # Data models/entities
│   ├── repository/         # Data access layer
│   └── service/            # Business logic
├── pkg/                    # Public packages (can be imported by other projects)
│   └── logger/             # Structured logging
├── scripts/                # Build and deployment scripts
├── Dockerfile              # Container build definition
├── docker-compose.yml      # Local development setup
├── Makefile                # Build automation
├── go.mod                  # Go module definition
└── main.go                 # Application entry point
```

## Prerequisites

- Go 1.22 or higher
- Docker (optional, for containerized deployment)
- Make (optional, for build automation)

## Getting Started

### Installation

```bash
# Clone the repository
git clone https://github.com/Aga602/github-slack-integration.git
cd github-slack-integration

# Download dependencies
go mod download

# Build the application
make build
# or
go build -o bin/github-slack-integration .
```

### Configuration

Configuration is managed through environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `ENVIRONMENT` | Environment name | `development` |
| `LOG_LEVEL` | Log level (debug, info, warn, error) | `info` |
| `SLACK_WEBHOOK_URL` | Slack webhook URL | - |

Copy `configs/config.example.env` to `.env` and update the values as needed.

### Running the Application

```bash
# Run directly
make run
# or
go run .

# Run with Docker
make docker-build
make docker-run

# Run with docker-compose
docker-compose up -d
```

## Development

### Available Make Commands

```bash
make help           # Show available commands
make build          # Build the application
make test           # Run tests
make coverage       # Run tests with coverage report
make lint           # Run linters
make fmt            # Format code
make tidy           # Tidy dependencies
make clean          # Clean build artifacts
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run tests directly
go test -v -race ./...
```

### Code Quality

This project uses [golangci-lint](https://golangci-lint.run/) for linting. Install it with:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Run the linter:

```bash
make lint
```

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/health` | GET | Health check endpoint |
| `/ready` | GET | Readiness check endpoint |
| `/api/v1/` | GET | API welcome message |

## Docker

### Build Docker Image

```bash
docker build -t github-slack-integration:latest .
```

### Run Docker Container

```bash
docker run -p 8080:8080 \
  -e PORT=8080 \
  -e LOG_LEVEL=info \
  github-slack-integration:latest
```

## CI/CD

This project includes GitHub Actions workflows for:

- **Build and Test**: Runs on every push and pull request
- **Security Scan**: Runs security analysis with gosec
- **Deployment**: Manual deployment workflow

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
