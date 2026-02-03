package main

import (
	"flag"
	"fmt"
	"os"
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

	fmt.Printf("Loading SQL script from %s into %s database\n", *scriptFile, *driver)
	// TODO: Implement actual script loading and execution
	return nil
}
