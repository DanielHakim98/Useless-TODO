package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type ServerAPI struct {
	DB *pgx.Conn
}

// (GET /todos)
func (server ServerAPI) FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := server.DB.Query(r.Context(),
		`
		SELECT
			id,
			title,
			COALESCE(to_char(updated_at, 'MM-DD-YYYY HH24:MI:SS'), '') AS date,
			content
		FROM useless_todo.todo_list`)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		errRes := Error{
			Code:    http.StatusInternalServerError,
			Message: "Error occured during DB query",
		}
		json.NewEncoder(w).Encode(errRes)
		return
	}
	todoList := []Todo{}
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Date, &todo.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		todoList = append(todoList, todo)
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
