package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	api := ServerAPI{}
	r := chi.NewRouter()
	r.Mount("/api/v1/", Handler(api))
	fmt.Println("Running server")
	http.ListenAndServe(":8080", r)
}
