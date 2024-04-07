package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/DanielHakim98/Useless-TODO/server"
	"github.com/go-chi/chi/v5"
)

func initRoute(server api.ServerInterface) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/api/v1/", api.Handler(server))
	return r
}

func main() {
	conn, err := server.GetDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database after multiple retries: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	sDB := server.ServerDB{Core: conn}
	r := initRoute(server.ServerAPI{DB: sDB})

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
