package main

import (
	"fmt"
	"net/http"

	"github.com/DanielHakim98/Useless-TODO/uselessTodoSpec"
	"github.com/go-chi/chi/v5"
)

func main() {
	api := uselessTodoSpec.TodoAPI{}
	r := chi.NewRouter()
	r.Mount("/useless-todo", uselessTodoSpec.Handler(&api))
	fmt.Println("Running server")
	http.ListenAndServe(":8080", r)
}
