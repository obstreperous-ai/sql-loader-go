package database

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name    string
		dsn     string
		wantErr bool
	}{
		{
			name:    "valid sqlite connection",
			dsn:     ":memory:",
			wantErr: false,
		},
		{
			name:    "invalid connection string",
			dsn:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := Connect("sqlite", tt.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && db != nil {
				defer db.Close()
				if err := db.Ping(); err != nil {
					t.Errorf("Ping() error = %v", err)
				}
			}
		})
	}
}

func TestExecuteScript(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	tests := []struct {
		name    string
		script  string
		wantErr bool
	}{
		{
			name: "create table",
			script: `CREATE TABLE users (
				id INTEGER PRIMARY KEY,
				name TEXT NOT NULL
			);`,
			wantErr: false,
		},
		{
			name: "insert data",
			script: `INSERT INTO users (name) VALUES ('Alice');
			INSERT INTO users (name) VALUES ('Bob');`,
			wantErr: false,
		},
		{
			name:    "invalid SQL",
			script:  "INVALID SQL STATEMENT",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ExecuteScript(db, tt.script)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteScript() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
