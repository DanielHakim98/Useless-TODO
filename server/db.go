package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/jackc/pgx/v5"
)

var (
	dbOnce sync.Once
)

func GetDB(cfg DBConfig) (conn *pgx.Conn, err error) {
	dbOnce.Do(func() {
		conn, err = initDB(cfg)
	})
	return
}

func initDB(cfg DBConfig) (conn *pgx.Conn, err error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.dbUser, cfg.dbPassword, cfg.dbHostname, cfg.dbPort, cfg.dbName)

	retry := 0
	maxRetry := 5
	for retry < maxRetry {
		conn, err = pgx.Connect(context.Background(), connString)
		if err == nil {
			return conn, nil
		}
		fmt.Fprintf(os.Stderr, "Unable to connect to database. Retrying in 5 seconds. Error: %v\n", err)
		time.Sleep(5 * time.Second)
	}

	return nil, err
}

type ServerDB struct {
	Core *pgx.Conn
}

func (sdb ServerDB) FindTodos(ctx context.Context, todoList *[]api.Todo) (err error) {
	rows, err := sdb.Core.Query(
		ctx,
		`
		SELECT
			id,
			title,
			content,
			to_char(created_at at time zone 'UTC','YYYY-MM-DD"T"HH24:MI:SS"Z"')
		FROM todo_list`)
	if err != nil {
		return err
	}
	for rows.Next() {
		todo := api.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Date)
		if err != nil {
			log.Println(err)
			continue
		}
		*todoList = append(*todoList, todo)
	}
	return nil
}

func (sdb ServerDB) AddTodo(ctx context.Context, body api.AddTodoJSONRequestBody) (api.Todo, error) {
	rows := sdb.Core.QueryRow(
		ctx,
		`	INSERT INTO todo_list (title, content, created_at)
			VALUES ($1, $2, now())
			RETURNING id, to_char(created_at at time zone 'UTC','YYYY-MM-DD"T"HH24:MI:SS"Z"')
		`, body.Title, body.Content)

	var todo api.Todo
	todo.Title = body.Title
	todo.Content = body.Content
	err := rows.Scan(&todo.Id, &todo.Date)
	if err != nil {
		return api.Todo{}, err
	}

	return todo, nil
}

func (sdb ServerDB) DeleteTodo(ctx context.Context, id int64) (api.Todo, error) {
	rows := sdb.Core.QueryRow(
		ctx,
		` DELETE FROM todo_list
		  WHERE id = $1
		  RETURNING
			title,
			content,
			to_char(created_at at time zone 'UTC','YYYY-MM-DD"T"HH24:MI:SS"Z"')
		`, id,
	)

	var todo api.Todo
	todo.Id = id
	err := rows.Scan(&todo.Title, &todo.Content, &todo.Date)
	if err != nil {
		return api.Todo{}, err
	}

	return todo, nil
}

func (sdb ServerDB) FindTodoById(ctx context.Context, id int64) (api.Todo, error) {
	rows := sdb.Core.QueryRow(
		ctx,
		`
		SELECT
			id,
			title,
			content,
			to_char(created_at at time zone 'UTC','YYYY-MM-DD"T"HH24:MI:SS"Z"')
		FROM todo_list
		WHERE id=$1
		`, id)
	var todo api.Todo
	err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Date)
	if err != nil {
		log.Println(err)
		return api.Todo{}, err
	}

	return todo, nil
}
