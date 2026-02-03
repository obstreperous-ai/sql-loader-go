package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// Connect establishes a database connection with the specified driver and DSN.
func Connect(driver, dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("DSN cannot be empty")
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// ExecuteScript executes a SQL script, splitting by semicolons and executing each statement.
func ExecuteScript(db *sql.DB, script string) error {
	script = strings.TrimSpace(script)
	if script == "" {
		return nil
	}

	statements := strings.Split(script, ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	return nil
}
