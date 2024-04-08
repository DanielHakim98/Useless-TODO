package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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
		FROM todo_list
		WHERE
			deleted_at IS NULL
		`)
	if err != nil {
		return err
	}
	defer rows.Close()

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
	// This adding todo method doesn't take into account
	// for identical 'title' and 'content' data from soft-deleted.
	// So even when once record soft-deleted and new identical record
	// is added, the record is treated as new record without updating
	// previous soft-delted record with exact same 'content' and 'title'

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
		`	UPDATE todo_list
			SET deleted_at=now()
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
		WHERE
			id=$1
			AND deleted_at IS NULL
		`, id)
	var todo api.Todo
	err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Date)
	if err != nil {
		log.Println(err)
		return api.Todo{}, err
	}

	return todo, nil
}

func (sdb ServerDB) UpdateTodoById(ctx context.Context,
	id int64, body api.NewTodo) (api.Todo, error) {

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(`
		UPDATE todo_list
		SET updated_at = now()
	`)

	var queryParams []interface{}
	paramIndex := 1

	if len(body.Title) > 0 {
		queryBuffer.WriteString(", title = $" + strconv.Itoa(paramIndex))
		queryParams = append(queryParams, body.Title)
		paramIndex++
	}

	if len(body.Content) > 0 {
		queryBuffer.WriteString(", content = $" + strconv.Itoa(paramIndex))
		queryParams = append(queryParams, body.Content)
		paramIndex++
	}

	queryParams = append(queryParams, id)

	queryBuffer.WriteString(
		" WHERE id = $" + strconv.Itoa(paramIndex) +
			`
				AND deleted_at IS NULL
				RETURNING
					id,
					title,
					content,
					to_char(created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"')`)
	query := queryBuffer.String()

	rows := sdb.Core.QueryRow(
		ctx, query, queryParams...)
	var todo api.Todo
	err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Date)
	if err != nil {
		log.Println(err)
		return api.Todo{}, err
	}

	return todo, nil

}
