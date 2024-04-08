package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/jackc/pgx/v5"
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

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// (DELETE /todos/{id})
func (server ServerAPI) DeleteTodo(w http.ResponseWriter, r *http.Request, id int64) {
	w.Header().Set("Content-Type", "application/json")
	deletedTodo, err := server.DB.DeleteTodo(r.Context(), id)
	if err != nil {
		log.Println(err)

		// ignore non-existing id from DB. Treat it as if it's deleted
		if err == pgx.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		server.errorResponse(w, http.StatusInternalServerError,
			"Error occured while deleting DB data")
		return
	}

	// For debugging purpose only, not needed for actual live environment
	by, _ := json.Marshal(deletedTodo)
	fmt.Fprintln(os.Stderr, string(by))

	w.WriteHeader(http.StatusNoContent)
}

// (GET /todos/{id})
func (server ServerAPI) FindTodoById(w http.ResponseWriter, r *http.Request, id int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
