// Package loader provides functionality for loading SQL scripts from files.
// It handles file I/O operations needed to read SQL scripts for execution.
package loader

import (
	"fmt"
	"os"
)

// LoadScript reads a SQL script file and returns its contents.
// The path parameter is expected to be a user-provided file path.
// #nosec G304 -- File path is intentionally provided by the user as part of the CLI interface
func LoadScript(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("script path cannot be empty")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read script file: %w", err)
	}

	return string(content), nil
}
