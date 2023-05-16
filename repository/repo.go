package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/todo-list/model"
)

type Repository interface {
	CreateToDo(ctx context.Context, tx *sql.Tx, category model.Course) model.Course
	UpdateToDo(ctx context.Context, tx *sql.Tx, category model.Course) model.Course
	DeleteToDo(ctx context.Context, tx *sql.Tx, id string)
	FindAllToDo(ctx context.Context, tx *sql.Tx) []model.Course
	FindTodo(ctx context.Context, tx *sql.Tx, id string) (model.Course, error)
	CreateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity
	UpdateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity
	DeleteActivity(ctx context.Context, tx *sql.Tx, id int)
	FindAllActivity(ctx context.Context, tx *sql.Tx) []model.Activity
	FindActivityById(ctx context.Context, tx *sql.Tx, id int) (model.Activity, error)
}
