package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/DanielHakim98/Useless-TODO/db"
)

type ServerAPI struct {
	DB db.ServerDB
}

// (GET /todos)
func (server ServerAPI) FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todoList := []api.Todo{}
	err := server.DB.FindTodos(r.Context(), &todoList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		errRes := api.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error occured during DB query",
		}
		json.NewEncoder(w).Encode(errRes)
		return
	}

	by, err := json.Marshal(todoList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		errRes := api.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while preparing JSON response",
		}
		json.NewEncoder(w).Encode(errRes)
		return
	}
	w.Write(by)
}

// (POST /todos)
func (server ServerAPI) AddTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /todos/{id})
func (server ServerAPI) DeleteTodo(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /todos/{id})
func (server ServerAPI) FindTodoById(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
