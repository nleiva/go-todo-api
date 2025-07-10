# Go TODO API

A modern, fast, and secure REST API for managing TODO items, built with Go, Fiber, GORM, and HTMX. Features a beautiful web interface, JWT authentication, and support for both SQLite and MySQL databases.

*Based on [TKSpectro/go-todo-api](https://github.com/TKSpectro/go-todo-api)*

## Features

- **High Performance**: Built with Go and Fiber for lightning-fast responses
- **Secure Authentication**: JWT-based authentication with bcrypt password hashing
- **Modern UI**: Beautiful, responsive web interface using HTMX and Tailwind CSS
- **Interactive Experience**: Real-time updates without page refreshes
- **Flexible Database**: Support for both SQLite (development) and MySQL (production)
- **Well Tested**: Comprehensive test suite using Go's standard testing package
- **API Documentation**: Auto-generated Swagger documentation
- **Developer Friendly**: Hot reloading, easy setup, and comprehensive tooling

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

**Note**: For development and testing, SQLite is used by default and requires no additional configuration.

2. **Run database migrations** (skip this step if using SQLite)
   ```bash
   make migrate-up
   ```

3. **Install dependencies and generate templates**
   ```bash
   go mod download
   go install github.com/a-h/templ/cmd/templ@latest
   templ generate
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

## Web Interface

Once the server is running, you can access:

- **Home Page**: `http://localhost:3000/` - Welcome page with features overview
- **Register**: `http://localhost:3000/register` - Create a new account
- **Login**: `http://localhost:3000/login` - Sign in to your account
- **Todos**: `http://localhost:3000/todos` - Manage your todo items (after login)
- **API Documentation**: `http://localhost:3000/swagger` - Interactive API documentation

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
| `make migrate-up` | Apply database migrations |
| `make migrate-down` | Rollback database migrations |
| `make docker-up` | Start Docker services (MySQL) |
| `make docker-down` | Stop Docker services |

### Project Structure

```
├── api/                    # API documentation (Swagger)
├── config/                 # Configuration management
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

This project supports both SQLite and MySQL databases with automatic schema generation.

### Environment-Based Database Selection

The database type is determined by the `DB_TYPE` environment variable:

```bash
# For SQLite (default for development/testing)
DB_TYPE=sqlite

# For MySQL (recommended for production)
DB_TYPE=mysql
```

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

# Check migration status
atlas migrate status --env gorm
```

## Technical Details

### Template Engine (TEMPL)
This project uses [TEMPL](https://templ.guide/) for type-safe HTML templates that compile to Go code.

**Key Benefits:**
- Type safety at compile time
- Fast rendering performance
- IntelliSense support in editors
- Hot reloading during development

Generate template files after making changes:
```bash
templ generate
# or using the Make command
make templ-generate
```

### Frontend Technologies
- **HTMX**: For dynamic, interactive web pages without JavaScript frameworks
- **Tailwind CSS**: For modern, responsive styling
- **Custom Components**: Reusable UI components built with TEMPL

### JSON Handling
The API uses custom null-handling for proper JSON serialization:

- `zero.String` - for optional string fields that can be empty
- `null.String` - for nullable string fields

Add `swaggertype:"string"` tags for proper API documentation.

### Authentication & Security
- **JWT-based authentication** with configurable token expiration
- **bcrypt password hashing** for secure password storage
- **Role-based permissions system** for access control
- **Secure session management** with automatic token refresh
- **CORS middleware** for cross-origin request handling

### API Features
- **RESTful endpoints** following OpenAPI 3.0 specification
- **Auto-generated Swagger documentation** available at `/swagger`
- **Request validation** using struct tags and custom validators
- **Error handling** with consistent JSON error responses
- **Pagination support** for large datasets

## Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Make your changes** and ensure they follow the project standards
4. **Run tests** (`make test`) to ensure everything works
5. **Update documentation** if needed
6. **Commit your changes** (`git commit -m 'Add amazing feature'`)
7. **Push to the branch** (`git push origin feature/amazing-feature`)
8. **Open a Pull Request** with a clear description of your changes

### Development Guidelines

- Follow Go best practices and conventions
- Write tests for new functionality
- Update documentation for API changes
- Use conventional commit messages
- Ensure all tests pass before submitting

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Original project by [TKSpectro](https://github.com/TKSpectro/go-todo-api)
- Built with amazing open-source libraries:
  - [Fiber](https://gofiber.io/) - Express-inspired web framework
  - [GORM](https://gorm.io/) - The fantastic ORM library for Go
  - [TEMPL](https://templ.guide/) - Type-safe HTML templates
  - [HTMX](https://htmx.org/) - Modern web interactions
  - [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
