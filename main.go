package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func initDB() (conn *pgx.Conn, err error) {
	retry := 0
	maxRetry := 5
	for retry < maxRetry {
		conn, err = pgx.Connect(context.Background(), "postgres://postgres:password@db:5432/postgres")
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

	r := initRoute(ServerAPI{DB: conn})
	fmt.Println("Running server")
	http.ListenAndServe(":8080", r)
}
