// Package main is the entry point for the sql-loader CLI application.
// It provides a command-line interface to execute SQL scripts against databases.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/obstreperous-ai/sql-loader-go/internal/database"
	"github.com/obstreperous-ai/sql-loader-go/internal/loader"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		showVersion = flag.Bool("version", false, "Show version information")
		driver      = flag.String("driver", "postgres", "Database driver (postgres, sqlite)")
		dsn         = flag.String("dsn", "", "Database connection string")
		scriptFile  = flag.String("file", "", "SQL script file to execute")
	)

	flag.Parse()

	if *showVersion {
		fmt.Printf("sql-loader version %s (commit: %s, built: %s)\n", version, commit, date)
		return nil
	}

	if *dsn == "" {
		return fmt.Errorf("DSN is required (use -dsn flag)")
	}

	if *scriptFile == "" {
		return fmt.Errorf("script file is required (use -file flag)")
	}

	// Load SQL script from file
	script, err := loader.LoadScript(*scriptFile)
	if err != nil {
		return fmt.Errorf("failed to load script: %w", err)
	}

	// Connect to database
	db, err := database.Connect(*driver, *dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close database: %v\n", closeErr)
		}
	}()

	// Execute script
	fmt.Printf("Loading SQL script from %s into %s database\n", *scriptFile, *driver)
	if err := database.ExecuteScript(db, script); err != nil {
		return fmt.Errorf("failed to execute script: %w", err)
	}

	fmt.Println("Script executed successfully")
	return nil
}
