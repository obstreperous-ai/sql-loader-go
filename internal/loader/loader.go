package loader

import (
	"fmt"
	"os"
)

// LoadScript reads a SQL script file and returns its contents.
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
