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

### Mandatory Pre-Completion Quality Checks

**CRITICAL**: Before reporting completion or finalizing any task, you MUST run ALL of these checks and fix any issues:

1. **Format Code**: `go fmt ./...` (must show no changes)
2. **Run Tests**: `go test -v -race ./...` (all tests must pass)
3. **Build Binary**: `go build -o sql-loader ./cmd/sql-loader` (must succeed)
4. **Lint Code**: `golangci-lint run ./...` (must have zero errors/warnings)
5. **Verify Binary**: `./sql-loader -version` (must run successfully)

**These checks are NOT optional**. If any check fails, you must fix the issues before proceeding. CI/CD will run these same checks and will fail if they don't pass.

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
7. **Zero Tolerance for Failures**: All tests, builds, and lints must pass before completion

## CI/CD Integration

### Understanding CI/CD Failures

When CI/CD failures are mentioned:

1. **Check Workflow Runs**: Use GitHub Actions tools to view recent workflow runs
2. **Review Job Logs**: Get detailed logs for failed jobs to understand what went wrong
3. **Common Failure Patterns**:
   - **Lint Failures**: Usually due to code style issues or golangci-lint version mismatch
   - **Build Failures**: Missing dependencies, syntax errors, or incompatible Go version
   - **Test Failures**: Logic errors, race conditions, or environmental issues

### Preventing CI/CD Failures

**Always run the Mandatory Pre-Completion Quality Checks locally before finalizing work.** This prevents CI/CD failures by catching issues early.

If golangci-lint is not available locally:
```bash
# Install golangci-lint (version 1.64+ for Go 1.24+)
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.64.0/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.0
```

## What to Modify

✅ **Allowed**:
- Source code in `cmd/` and `internal/` directories
- Test files (`*_test.go`)
- Documentation files (`*.md`)
- Build and CI configuration (`.github/workflows/`, `Taskfile.yml`, `.golangci.yml`, etc.)
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
- **CRITICAL**: Run ALL Mandatory Pre-Completion Quality Checks before submitting
- Ensure all tests pass before submitting
- Ensure golangci-lint runs with zero errors/warnings
- Follow the project's minimalism principle
- Never submit a PR with failing CI/CD checks
