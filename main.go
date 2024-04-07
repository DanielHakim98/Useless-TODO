package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/DanielHakim98/Useless-TODO/db"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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

func initDB() (conn *pgx.Conn, err error) {
	cfg := getConfig()
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.dbUser, cfg.dbPassword, cfg.dbHostname, cfg.dbPort, cfg.dbName)

	retry := 0
	maxRetry := 5
	for retry < maxRetry {
		conn, err = pgx.Connect(context.Background(), connString)
		if err == nil {
			return conn, nil
		}
		fmt.Fprintf(os.Stderr, "Unable to connect to database. Retrying in 5 seconds. Error: %v\n", err)
		time.Sleep(5 * time.Second)
	}

	return nil, err
}

func initRoute(server api.ServerInterface) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/api/v1/", api.Handler(server))
	return r
}

func main() {
	conn, err := initDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database after multiple retries: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	sDB := db.ServerDB{Core: conn}
	r := initRoute(ServerAPI{DB: sDB})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80"
	}

	hostname := os.Getenv("SERVER_HOSTNAME")
	if hostname == "" {
		hostname = ""
	}

	fmt.Fprintln(os.Stderr, "Running server at '"+hostname+":"+port+"'")
	http.ListenAndServe(hostname+":"+port, r)
}
