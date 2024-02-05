package main

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type ServerAPI struct {
	DB *pgx.Conn
}

// (GET /todos)
func (server ServerAPI) FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todoList []Todo
	err := server.DB.QueryRow(
		r.Context(),
		`SELECT id, title, date, content FROM useless_todo.todo_list`).Scan(&todoList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errRes := Error{
			Code:    http.StatusInternalServerError,
			Message: "Error occured during DB query",
		}
		json.NewEncoder(w).Encode(errRes)
		return
	}

	by, _ := json.Marshal(todoList)
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
