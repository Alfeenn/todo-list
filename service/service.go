package service

import (
	"context"

	"github.com/Alfeenn/todo-list/model"
)

type Service interface {
	CreateToDo(ctx context.Context, req model.Course) model.Course
	UpdateToDo(ctx context.Context, req model.Course) model.Course
	DeleteToDo(ctx context.Context, id string)
	FindTodo(ctx context.Context, id string) model.Course
	FindAllToDo(ctx context.Context) []model.Course
	CreateActivity(ctx context.Context, request model.Activity) model.Activity
	UpdateActivity(ctx context.Context, request model.Activity) model.Activity
	DeleteActivity(ctx context.Context, id int)
	FindAllActivity(ctx context.Context) []model.Activity
	FindActivityById(ctx context.Context, id int) model.Activity
}
