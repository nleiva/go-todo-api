# Go TODO API

A REST API for managing TODO items built with Go, Fiber, and GORM.

*Based on [TKSpectro/go-todo-api](https://github.com/TKSpectro/go-todo-api)*

## Quick Start

### Prerequisites

- [Go](https://golang.org/) 1.23 or higher
- [Docker](https://www.docker.com/) (optional - for MySQL database)
- [Make](https://www.gnu.org/software/make/) (optional - for running shortcuts)
- [Air](https://github.com/cosmtrek/air/) (optional - for hot reloading during development)

### Installation

1. **Clone and setup environment**
   ```bash
   git clone <repository-url>
   cd go-todo-api
   cp .env.example .env
   ```
   Edit the `.env` file with your configuration values.

2. **Run database migrations** (skip this step if using SQLite)
   ```bash
   make migrate-up
   ```

3. **Generate template files**
   ```bash
   go get -tool github.com/a-h/templ/cmd/templ@latest
   go tool templ generate
   ```

4. **Start the server**
   ```bash
   make build-run
   # or for development with hot reload
   air
   ```

5. **Verify installation** (optional)
   ```bash
   make test
   ```

The API will be available at `http://localhost:3000`

## Development

### Available Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the application |
| `make run` | Run the built application |
| `make build-run` | Build and run in one command |
| `make test` | Run all tests |
| `make test-v` | Run tests with verbose output |
| `make test-coverage` | Run tests with coverage report |
| `make docker-up` | Start Docker services |
| `make docker-down` | Stop Docker services |

### Project Structure

```
├── api/                    # API documentation (Swagger)
├── config/                 # Configuration management
├── migrations/             # Database migration files
├── pkg/
│   ├── app/
│   │   ├── handler/        # HTTP route handlers
│   │   ├── model/          # Data models
│   │   ├── service/        # Business logic
│   │   └── types/          # Type definitions
│   ├── database/           # Database connection logic
│   ├── jwt/                # JWT authentication
│   ├── middleware/         # HTTP middleware
│   └── view/               # HTML templates
├── test/                   # Test utilities
└── utils/                  # Helper functions
```

## Testing

This project uses Go's standard testing framework for reliability and simplicity.

### Running Tests

```bash
# Run all tests
make test

# Run with verbose output
make test-v

# Run with coverage report
make test-coverage

# Run for CI/CD
make test-ci
```

### Test Organization

- **Model Tests** (`pkg/app/model/*_test.go`) - Unit tests for data models and business logic
- **Handler Tests** (`pkg/app/handler/*_test.go`) - Integration tests for HTTP endpoints

**Key Features:**
- Uses SQLite in-memory database for fast, isolated test execution
- No external dependencies required for testing
- Automatic setup and teardown between test runs
- Previously migrated from Ginkgo to standard Go testing for better tooling integration

## Database Configuration

### SQLite (Default/Testing)
- **Usage**: Testing and development
- **Configuration**: In-memory database, automatically configured
- **Migrations**: Applied automatically during startup

### MySQL (Production)
- **Usage**: Production deployments
- **Configuration**: Set database credentials in `.env` file
- **Migrations**: Managed through Atlas migration tool

### Migration Commands

```bash
# Install Atlas migration tool
curl -sSf https://atlasgo.sh | sh

# Generate migration from model changes
make migrate-gen name=<migration-name>

# Create empty migration file
make migrate-new name=<migration-name>

# Apply all pending migrations
make migrate-up

# Rollback to specific version
make migrate-down version=<version-timestamp>
```

## Technical Details

### Template Engine
This project uses [TEMPL](https://templ.guide/) for type-safe HTML templates.

Generate template files after making changes:
```bash
go tool templ generate
```

### JSON Handling
The API uses custom null-handling for proper JSON serialization:

- `zero.String` - for optional string fields that can be empty
- `null.String` - for nullable string fields

Add `swaggertype:"string"` tags for proper API documentation.

### Authentication
- JWT-based authentication
- Configurable token expiration
- Role-based permissions system

### API Documentation
Swagger documentation is auto-generated and available at `/swagger` when running the server.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`make test`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
