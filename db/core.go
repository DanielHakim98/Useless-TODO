package db

import (
	"context"
	"log"

	"github.com/DanielHakim98/Useless-TODO/api"
	"github.com/jackc/pgx/v5"
)

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
