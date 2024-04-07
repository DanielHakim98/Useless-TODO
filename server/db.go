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
