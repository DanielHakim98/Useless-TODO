// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NewTodo defines model for NewTodo.
type NewTodo struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}

// Todo defines model for Todo.
type Todo struct {
	Content string `json:"content"`
	Date    string `json:"date"`
	Id      int64  `json:"id"`
	Title   string `json:"title"`
}

// Todos defines model for Todos.
type Todos struct {
	Content string `json:"content"`
	Date    string `json:"date"`
	Id      int64  `json:"id"`
	Title   string `json:"title"`
}

// AddTodoJSONRequestBody defines body for AddTodo for application/json ContentType.
type AddTodoJSONRequestBody = NewTodo
