package loader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadScript(t *testing.T) {
	tmpDir := t.TempDir()

	validScript := filepath.Join(tmpDir, "valid.sql")
	if err := os.WriteFile(validScript, []byte("SELECT 1;"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name    string
		path    string
		want    string
		wantErr bool
	}{
		{
			name:    "valid script file",
			path:    validScript,
			want:    "SELECT 1;",
			wantErr: false,
		},
		{
			name:    "non-existent file",
			path:    filepath.Join(tmpDir, "missing.sql"),
			wantErr: true,
		},
		{
			name:    "empty path",
			path:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadScript(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("LoadScript() = %v, want %v", got, tt.want)
			}
		})
	}
}
