package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/go-chi/chi/v5"
)

func Start() {
	cfg := GetConfig()
	conn, err := GetDB(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database after multiple retries: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	sDB := ServerDB{Core: conn}
	sAPI := ServerAPI{DB: sDB}

	r := chi.NewRouter()
	r.Mount("/api/v1/", api.Handler(sAPI))

	fmt.Fprintln(os.Stderr, "Running server at '"+cfg.hostname+":"+cfg.port+"'")
	http.ListenAndServe(cfg.hostname+":"+cfg.port, r)
}
