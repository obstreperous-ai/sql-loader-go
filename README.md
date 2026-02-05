# sql-loader-go

[![Build and Test](https://github.com/obstreperous-ai/sql-loader-go/actions/workflows/build.yml/badge.svg)](https://github.com/obstreperous-ai/sql-loader-go/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/obstreperous-ai/sql-loader-go)](https://goreportcard.com/report/github.com/obstreperous-ai/sql-loader-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lean CLI utility in Go to execute SQL data load scripts in a locked-down container (single binary). Designed for minimalism, test-first development, and quality.

## Features

- **Single Binary**: Compiles to a single static binary with no external dependencies
- **Multiple Databases**: Support for PostgreSQL and SQLite
- **Containerized**: Designed for locked-down container environments
- **Minimal Dependencies**: Only essential dependencies, no bloat
- **Test-First**: Comprehensive test coverage with embedded database testing
- **Cross-Platform**: Builds for Linux, macOS, and Windows

## Installation

### Download Pre-built Binaries

Download the latest release from the [releases page](https://github.com/obstreperous-ai/sql-loader-go/releases).

```bash
# Linux (amd64)
wget https://github.com/obstreperous-ai/sql-loader-go/releases/latest/download/sql-loader-linux-amd64
chmod +x sql-loader-linux-amd64
mv sql-loader-linux-amd64 /usr/local/bin/sql-loader

# macOS (arm64/Apple Silicon)
wget https://github.com/obstreperous-ai/sql-loader-go/releases/latest/download/sql-loader-darwin-arm64
chmod +x sql-loader-darwin-arm64
mv sql-loader-darwin-arm64 /usr/local/bin/sql-loader
```

### Build from Source

Requirements:
- Go 1.23 or later

```bash
# Clone the repository
git clone https://github.com/obstreperous-ai/sql-loader-go.git
cd sql-loader-go

# Build
go build -o sql-loader ./cmd/sql-loader

# Or install to $GOPATH/bin
go install ./cmd/sql-loader
```

## Usage

```bash
# Show version
sql-loader -version

# Execute SQL script (PostgreSQL)
sql-loader -driver postgres -dsn "postgres://user:pass@localhost/db" -file script.sql

# Execute SQL script (SQLite)
sql-loader -driver sqlite -dsn "path/to/database.db" -file script.sql
```

### Flags

- `-driver`: Database driver (postgres, sqlite) [default: postgres]
- `-dsn`: Database connection string (required)
- `-file`: SQL script file to execute (required)
- `-version`: Show version information

### Example SQL Script

```sql
-- script.sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email) VALUES 
    ('Alice', 'alice@example.com'),
    ('Bob', 'bob@example.com');
```

## Development

### Prerequisites

- Go 1.23+
- [Task](https://taskfile.dev/) (optional, for build automation)
- golangci-lint (optional, for linting)

### Development with VS Code DevContainer

This project includes a DevContainer configuration for consistent development environments:

1. Install [VS Code](https://code.visualstudio.com/) and the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
2. Open the project in VS Code
3. Click "Reopen in Container" when prompted

The DevContainer includes Go 1.23 and all necessary tools.

### Building

```bash
# Using Task (recommended)
task build

# Using go directly
go build -o sql-loader ./cmd/sql-loader
```

### Testing

```bash
# Run all tests
task test

# Or with go
go test -v ./...

# With coverage
task test-coverage
go test -v -cover ./...
```

### Linting

```bash
# Using Task
task lint

# Or directly
golangci-lint run ./...
```

### Formatting

```bash
# Using Task
task fmt

# Or with go
go fmt ./...
```

### Cross-Platform Builds

```bash
# Build for all platforms
task package

# Binaries will be in ./build/
# - sql-loader-linux-amd64
# - sql-loader-linux-arm64
# - sql-loader-darwin-amd64
# - sql-loader-darwin-arm64
# - sql-loader-windows-amd64.exe
```

## Contributing

Contributions are welcome! This project follows test-first development and minimalist principles.

### Guidelines

1. **Test-First**: Write tests before implementation
2. **Minimalism**: Keep code simple and focused
3. **Quality**: Follow Go best practices and idioms
4. **Documentation**: Document public APIs
5. **Security**: Use parameterized queries, validate inputs

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Write tests for your changes
4. Implement your changes
5. **Run all quality checks** (see checklist below)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Pre-PR Quality Checklist

**Before opening a Pull Request, ensure ALL of these checks pass:**

```bash
# 1. Format code
go fmt ./...

# 2. Run tests with race detection
go test -v -race ./...

# 3. Build the binary
go build -o sql-loader ./cmd/sql-loader

# 4. Verify binary works
./sql-loader -version

# 5. Run linter (requires golangci-lint)
golangci-lint run ./...
```

**All checks must pass with zero errors/warnings before submitting a PR.** This prevents CI/CD failures and ensures code quality.

To install golangci-lint if not already installed:
```bash
# Install golangci-lint v1.64+ (required for Go 1.24+)
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.64.0/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.0
```

### Code Style

- Follow standard Go conventions
- Use table-driven tests
- Keep functions small and focused
- Handle all errors explicitly
- Document exported functions

## Project Structure

```
.
├── cmd/
│   └── sql-loader/       # Main application entry point
├── internal/
│   ├── database/         # Database connection and execution
│   └── loader/           # SQL script file loading
├── .devcontainer/        # VS Code DevContainer configuration
├── .github/
│   ├── workflows/        # GitHub Actions CI/CD
│   └── dependabot.yml    # Dependency updates
├── Taskfile.yml          # Build automation tasks
├── go.mod                # Go module definition
└── README.md             # This file
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go and minimal dependencies
- Uses [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) for embedded SQLite testing
- Designed for containerized environments

## Support

- Report bugs and request features via [GitHub Issues](https://github.com/obstreperous-ai/sql-loader-go/issues)
- For questions and discussions, use [GitHub Discussions](https://github.com/obstreperous-ai/sql-loader-go/discussions)

