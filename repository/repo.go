package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/todo-list/model"
)

type Repository interface {
	CreateToDo(ctx context.Context, tx *sql.Tx, category model.Todo) model.Todo
	UpdateToDo(ctx context.Context, tx *sql.Tx, category model.Todo) model.Todo
	DeleteToDo(ctx context.Context, tx *sql.Tx, id int)
	FindAllToDo(ctx context.Context, tx *sql.Tx) []model.Todo
	FindTodo(ctx context.Context, tx *sql.Tx, id int) (model.Todo, error)
	CreateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity
	UpdateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity
	DeleteActivity(ctx context.Context, tx *sql.Tx, id int)
	FindAllActivity(ctx context.Context, tx *sql.Tx) []model.Activity
	FindActivityById(ctx context.Context, tx *sql.Tx, id int) (model.Activity, error)
}
