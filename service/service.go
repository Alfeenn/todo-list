package service

import (
	"context"

	"github.com/Alfeenn/todo-list/model"
	"github.com/Alfeenn/todo-list/model/web"
)

type Service interface {
	CreateToDo(ctx context.Context, req model.Course) model.Course
	UpdateToDo(ctx context.Context, req model.Course) model.Course
	DeleteToDo(ctx context.Context, id string)
	FindTodo(ctx context.Context, id string) model.Course
	FindCourseByCategory(ctx context.Context, id string) model.Course
	FindAllToDo(ctx context.Context) []model.Course
	Login(ctx context.Context, request web.RequestLogin) web.CatResp
	Register(ctx context.Context, request web.CategoryRequest) web.CatResp
	GetCourse(ctx context.Context, req model.Class, id string) model.Class
	DeleteUser(ctx context.Context, id string)
}
