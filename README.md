# Go TODO API
[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/nleiva/go-todo-api?quickstart=1)

A modern, fast, and secure REST API for managing TODO items, built with Go, Fiber, GORM, and HTMX. It features a beautiful web interface, JWT authentication, and support for both SQLite and MySQL databases.

*Based on [TKSpectro/go-todo-api](https://github.com/TKSpectro/go-todo-api)*

## Features

- **High Performance**: Built with Go and Fiber for fast responses
- **Secure Authentication**: JWT-based authentication with bcrypt password hashing
- **Modern UI**: Beautiful, responsive web interface using HTMX and Tailwind CSS
- **Interactive Experience**: Real-time updates without page refreshes
- **Flexible Database**: Support for both SQLite (development) and MySQL (production)
- **Comprehensive Testing**: Test suite using Go's standard testing package
- **API Documentation**: Auto-generated OpenAPI documentation with Swagger UI and Redoc
- **Developer Tools**: Hot reloading, easy setup, and comprehensive tooling

## Quick Start

### Prerequisites

- [Go](https://golang.org/) 1.23 or higher
- [Docker](https://www.docker.com/) (optional - for MySQL database)
- [Make](https://www.gnu.org/software/make/) (optional - for running shortcuts)
- [Air](https://github.com/cosmtrek/air/) (optional - for hot reloading during development)

### Installation

1. **Clone the repository and set up environment**
   ```bash
   git clone <repository-url>
   cd go-todo-api
   cp .env.example .env
   ```
   Edit the `.env` file with your configuration values.

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```bash
# Server Configuration
PORT=3000
HOST=localhost

# Database Configuration
DB_TYPE=sqlite                    # or "mysql" for production
DB_HOST=localhost                 # MySQL host (if using MySQL)
DB_PORT=3306                      # MySQL port (if using MySQL)
DB_USER=root                      # MySQL username (if using MySQL)
DB_PASSWORD=password              # MySQL password (if using MySQL)
DB_NAME=go_todo_api              # Database name (if using MySQL)

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here
JWT_EXPIRY=24h                    # Token expiration time

# CORS Configuration
CORS_ORIGINS=http://localhost:3000,http://127.0.0.1:3000
```

**Note**: For development and testing, SQLite is used by default and requires no additional setup.

2. **Install dependencies and generate templates**
   ```bash
   go mod download
   go get -tool github.com/swaggo/swag/cmd/swag@latest
   go get -tool github.com/a-h/templ/cmd/templ@latest
   go tool templ generate
   ```

3. **Start the server**
   ```bash
   make build-run
   # or for development with hot reload
   air
   ```

4. **Run tests** (optional)
   ```bash
   make test
   ```

The API will be available at `http://localhost:3000`

## Web Interface

Once the server is running, you can access:

- **Home Page**: `http://localhost:3000/` - Welcome page with features overview
- **Register**: `http://localhost:3000/register` - Create a new account
- **Login**: `http://localhost:3000/login` - Sign in to your account
- **Todos**: `http://localhost:3000/todos` - Manage your todo items (requires login)
- **API Documentation**: `http://localhost:3000/api/docs` - Interactive Swagger UI documentation
- **Redoc Documentation**: `http://localhost:3000/api/redoc` - Clean API documentation with Redoc

## Development

### Available Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the application |
| `make run` | Run the built application |
| `make build-run` | Build and run in one command |
| `make test` | Run all tests with SQLite |
| `make test-v` | Run tests with verbose output |
| `make test-coverage` | Run tests with coverage report |
| `make docs` | Generate API documentation |
| `make migrate-up` | Apply database migrations |
| `make migrate-down` | Rollback database migrations |
| `make docker-up` | Start Docker services (MySQL) |
| `make docker-down` | Stop Docker services |

### Project Structure

```
├── config/                 # Configuration management
├── docs/                   # Generated API documentation
├── loader/                 # Database initialization and schema generation
├── migrations/             # Database migration files
├── pkg/
│   ├── app/
│   │   ├── handler/        # HTTP route handlers
│   │   ├── model/          # Data models and database entities
│   │   ├── service/        # Business logic layer
│   │   └── types/          # Type definitions and DTOs
│   ├── database/           # Database connection logic (SQLite/MySQL)
│   ├── jwt/                # JWT token generation and validation
│   ├── jwk/                # JSON Web Key management
│   ├── middleware/         # HTTP middleware (auth, CORS, etc.)
│   ├── permission/         # Permission and authorization logic
│   └── view/               # TEMPL templates for web interface
├── test/                   # Test utilities and Docker setup
├── tmp/                    # Build artifacts
└── utils/                  # Helper functions and utilities
```

## Testing

This project uses Go's standard testing framework.

### Running Tests

```bash
# Run all tests
make test

# Run with verbose output
make test-v

# Run with coverage report
make test-coverage

# Run for CI/CD environments
make test-ci
```

### Test Organization

- **Model Tests** (`pkg/app/model/*_test.go`) - Unit tests for data models and business logic
- **Handler Tests** (`pkg/app/handler/*_test.go`) - Integration tests for HTTP endpoints

The tests use SQLite in-memory database for fast, isolated test execution with no external dependencies required.

## Database Configuration

This project supports both SQLite and MySQL databases.

### Database Selection

The database type is determined by the `DB_TYPE` environment variable:

```bash
# For SQLite (default for development/testing)
DB_TYPE=sqlite

# For MySQL (recommended for production)
DB_TYPE=mysql
```

### SQLite (Default)
- **Usage**: Development and testing
- **Configuration**: Automatically configured
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

# Check migration status
atlas migrate status --env gorm
```

## Technical Details

### Template Engine (TEMPL)
This project uses [TEMPL](https://templ.guide/) for type-safe HTML templates that compile to Go code.

Benefits include type safety at compile time, fast rendering performance, and IntelliSense support in editors.

Generate template files after making changes:
```bash
templ generate
```

### Frontend Technologies
- **HTMX**: For dynamic, interactive web pages without JavaScript frameworks
- **Tailwind CSS**: For modern, responsive styling
- **Custom Components**: Reusable UI components built with TEMPL

### JSON Handling
The API uses `zero.String` for optional string fields that can be empty, with `swaggertype:"string"` tags for proper API documentation.

### Authentication & Security
- **JWT-based authentication** with configurable token expiration
- **bcrypt password hashing** for secure password storage
- **Role-based permissions system** for access control
- **Secure session management** with automatic token refresh
- **CORS middleware** for cross-origin request handling

### API Features
- **RESTful endpoints** following OpenAPI 3.0 specification
- **Auto-generated documentation** available via Swagger UI and Redoc
- **Interactive API testing** with Swagger UI interface
- **Request validation** using struct tags and custom validators
- **Error handling** with consistent JSON error responses
- **Pagination support** for large datasets

Generate API documentation after making changes:
```bash
make docs
```

## Acknowledgments

- Original project by [TKSpectro](https://github.com/TKSpectro/go-todo-api)
- Built with excellent open-source libraries:
  - [Fiber](https://gofiber.io/) - Express-inspired web framework
  - [GORM](https://gorm.io/) - ORM library for Go
  - [TEMPL](https://templ.guide/) - Type-safe HTML templates
  - [HTMX](https://htmx.org/) - Modern web interactions
  - [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
