package main

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type ServerAPI struct {
	DB *gorm.Model
}

// (GET /todos)
func (server ServerAPI) FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todoList []Todo
	todoList = append(todoList, Todo{})
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
