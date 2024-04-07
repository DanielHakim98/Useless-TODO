package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/DanielHakim98/Useless-TODO/api"
)

type ServerAPI struct {
	DB ServerDB
}

func (server ServerAPI) errorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	errRes := api.Error{
		Code:    int32(statusCode),
		Message: errorMessage,
	}
	json.NewEncoder(w).Encode(errRes)
}

// (GET /todos)
func (server ServerAPI) FindTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todoList := []api.Todo{}
	err := server.DB.FindTodos(r.Context(), &todoList)
	if err != nil {
		log.Println(err)
		server.errorResponse(w, http.StatusInternalServerError,
			"Error occured while querying DB data")
		return
	}

	by, err := json.Marshal(todoList)
	if err != nil {
		log.Println(err)
		server.errorResponse(w, http.StatusInternalServerError,
			"Error occured while generating JSON response")
		return
	}
	w.Write(by)
}

func (server ServerAPI) validateNewTodo(body api.AddTodoJSONRequestBody) error {
	if strings.TrimSpace(body.Content) == "" {
		return fmt.Errorf("missing value key 'content'")
	}

	if strings.TrimSpace(body.Title) == "" {
		return fmt.Errorf("missing value key 'title'")
	}

	return nil
}

// (POST /todos)
func (server ServerAPI) AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	by, err := io.ReadAll(r.Body)
	if err != nil {
		server.errorResponse(w, http.StatusInternalServerError,
			"Error occured while parsing request body")
		return
	}

	var body api.AddTodoJSONRequestBody
	err = json.Unmarshal(by, &body)
	if err != nil {
		log.Println(err)
		server.errorResponse(w, http.StatusBadRequest,
			"Invalid request body format/structure")
		return
	}

	err = server.validateNewTodo(body)
	if err != nil {
		log.Println(err)
		server.errorResponse(w, http.StatusUnprocessableEntity,
			err.Error())
		return
	}

	todo, err := server.DB.AddTodo(r.Context(), body)
	if err != nil {
		log.Println(err)
		server.errorResponse(w, http.StatusInternalServerError,
			"Error occured while creating DB data")
		return
	}

	res, err := json.Marshal(todo)
	if err != nil {
		log.Println(err)
		server.errorResponse(w, http.StatusInternalServerError,
			"Error occured while generating JSON response")
		return
	}

	w.Write(res)
}

// (DELETE /todos/{id})
func (server ServerAPI) DeleteTodo(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /todos/{id})
func (server ServerAPI) FindTodoById(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
