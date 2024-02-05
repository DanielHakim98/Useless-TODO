package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func main() {
	var conn *pgx.Conn
	var err error

	for retry := 0; retry < 5; retry++ {
		conn, err = pgx.Connect(context.Background(),
			"postgres://postgres:password@db:5432/postgres")
		if err == nil {
			break
		}

		fmt.Fprintf(os.Stderr, "Unable to connect to database. Retrying in 5 seconds. Error: %v\n", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database after multiple retries: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
	api := ServerAPI{DB: conn}
	r := chi.NewRouter()
	r.Mount("/api/v1/", Handler(api))
	fmt.Println("Running server")
	http.ListenAndServe(":8080", r)
}
