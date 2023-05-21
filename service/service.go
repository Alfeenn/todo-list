package service

import (
	"context"

	"github.com/Alfeenn/todo-list/model"
)

type Service interface {
	CreateToDo(ctx context.Context, req model.Todo) model.Todo
	UpdateToDo(ctx context.Context, req model.Todo) model.Todo
	DeleteToDo(ctx context.Context, id int)
	FindTodo(ctx context.Context, id int) model.Todo
	FindAllToDo(ctx context.Context) []model.Todo
	CreateActivity(ctx context.Context, request model.Activity) model.Activity
	UpdateActivity(ctx context.Context, request model.Activity) model.Activity
	DeleteActivity(ctx context.Context, id int)
	FindAllActivity(ctx context.Context) []model.Activity
	FindActivityById(ctx context.Context, id int) model.Activity
}
