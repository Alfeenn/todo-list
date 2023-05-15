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
	FindCourseByCategory(ctx context.Context, tx *sql.Tx, category string) (model.Course, error)
	Login(ctx context.Context, tx *sql.Tx, category model.User) (model.User, error)
	Register(ctx context.Context, tx *sql.Tx, category model.User) model.User
	GetCourse(ctx context.Context, tx *sql.Tx, category model.Class, id string) model.Class
	DeleteCourse(ctx context.Context, tx *sql.Tx, id string)
}
