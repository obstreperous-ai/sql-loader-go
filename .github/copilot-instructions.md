# GitHub Copilot Instructions for sql-loader-go

## Project Overview

This is a lean CLI utility in Go to execute SQL data load scripts in a locked-down container (single binary). The project prioritizes:
- **Minimalism**: Keep the codebase small and focused
- **Test-first development**: Write tests before implementation
- **Quality**: High code quality standards and security

## Technology Stack

- **Language**: Go (latest stable version)
- **Database Testing**: 
  - Use embedded SQLite (`modernc.org/sqlite` or `github.com/mattn/go-sqlite3`)
  - Use pgx (`github.com/jackc/pgx/v5`) for PostgreSQL tests
  - **No external databases** - all tests must run in-process
- **Build**: Single static binary output
- **Packaging**: Designed for containerized environments

## Development Workflow

### Setup and Build Commands

```bash
# Build the project
go build -o sql-loader ./cmd/sql-loader

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Format code
go fmt ./...

# Lint code (if golangci-lint is available)
golangci-lint run
```

### Testing Standards

1. **Test-First Development**: Always write tests before implementation
2. **Table-Driven Tests**: Prefer table-driven test patterns for multiple test cases
3. **No External Dependencies**: Tests must not require external databases or services
4. **Embedded Databases**: Use in-memory SQLite or pgx with Docker container tests
5. **Coverage**: Aim for high test coverage on critical paths
6. **Fast Tests**: Tests should run quickly without external I/O when possible

### Code Quality Standards

1. **Minimalism**: Every line of code should have a clear purpose
2. **Go Idioms**: Follow standard Go conventions and idioms
3. **Error Handling**: Always handle errors explicitly, never ignore them
4. **Documentation**: Document public APIs with clear godoc comments
5. **Security**: Follow secure coding practices, especially for SQL execution
6. **Single Responsibility**: Keep functions and modules focused on one thing

## What to Modify

✅ **Allowed**:
- Source code in `cmd/` and `internal/` directories
- Test files (`*_test.go`)
- Documentation files (`*.md`)
- Build and CI configuration (`.github/workflows/`, `Makefile`, etc.)
- Go module files (`go.mod`, `go.sum`)

## What NOT to Modify

❌ **Prohibited**:
- The `.github/agents/` directory (agent configuration files)
- License file
- Git configuration files
- Dependencies without explicit justification

## Security Guidelines

1. **SQL Injection Prevention**: Use parameterized queries exclusively
2. **Input Validation**: Validate all user inputs
3. **Secure Defaults**: Default to secure configurations
4. **No Hardcoded Secrets**: Never commit credentials or secrets
5. **Minimal Dependencies**: Only add dependencies that are absolutely necessary

## Code Examples

### Test Example (Table-Driven)

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:  "valid input",
            input: "test",
            want:  "result",
        },
        {
            name:    "invalid input",
            input:   "",
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Function(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("Function() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Function() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Error Handling Example

```go
func DoSomething() error {
    if err := step1(); err != nil {
        return fmt.Errorf("step1 failed: %w", err)
    }
    if err := step2(); err != nil {
        return fmt.Errorf("step2 failed: %w", err)
    }
    return nil
}
```

## Commit Message Standards

- Use imperative mood: "Add feature" not "Added feature"
- Keep first line under 50 characters
- Add detailed explanation after blank line if needed
- Reference issue numbers when applicable

## Pull Request Standards

- Keep PRs small and focused
- Include tests for all new functionality
- Update documentation as needed
- Ensure all tests pass before submitting
- Follow the project's minimalism principle
