package server

import (
	"fmt"
	"os"
)

type DBConfig struct {
	dbUser, dbPassword, dbName, dbPort, dbHostname string
}

func getConfig() DBConfig {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Fprintln(os.Stderr, "Error: environment variable 'DB_USER' not found")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Fprintln(os.Stderr, "Error: environment variable 'DB_PASSWORD' not found")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Fprintln(os.Stderr, "Error: environment variable 'DB_NAME' not found")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Fprintln(os.Stderr, "Error: environment variable 'DB_PORT' not found")
		os.Exit(1)
	}

	dbHostname := os.Getenv("DB_HOSTNAME")
	if dbHostname == "" {
		fmt.Fprintln(os.Stderr, "Error: environment variable 'DB_HOSTNAME' not found")
		os.Exit(1)
	}

	return DBConfig{dbUser, dbPassword, dbName, dbPort, dbHostname}
}
