package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/Alfeenn/todo-list/helper"
	"github.com/Alfeenn/todo-list/middleware"
	"github.com/Alfeenn/todo-list/model"
	"github.com/Alfeenn/todo-list/repository"
)

type ServiceImpl struct {
	Rep repository.Repository
	DB  *sql.DB
}

func NewService(c repository.Repository, DB *sql.DB) Service {
	return &ServiceImpl{
		Rep: c,
		DB:  DB,
	}
}

func (s *ServiceImpl) CreateToDo(ctx context.Context, req model.Todo) model.Todo {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	request := model.Todo{
		Id:         req.Id,
		Title:      req.Title,
		ActivityId: req.ActivityId,
		Isactive:   req.Isactive,
		CreatedAt:  req.CreatedAt,
	}
	if request.Priority == "" {
		request.Priority = "very-high"
	}
	Todo := s.Rep.CreateToDo(ctx, tx, request)

	return Todo

}

func (s *ServiceImpl) UpdateToDo(ctx context.Context, req model.Todo) (model.Todo, error) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	id := req.Id
	model, err := s.Rep.FindTodo(ctx, tx, id)
	if err != nil {
		return model, err
	} else {

		model.Title = req.Title
		model = s.Rep.UpdateToDo(ctx, tx, model)
		return model, nil
	}
}

func (s *ServiceImpl) DeleteToDo(ctx context.Context, id int) error {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	req, err := s.Rep.FindTodo(ctx, tx, id)
	if err != nil {
		return err
	} else {
		s.Rep.DeleteToDo(ctx, tx, req.Id)
		return nil
	}

}

func (s *ServiceImpl) FindTodo(ctx context.Context, id int) model.Todo {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.FindTodo(ctx, tx, id)
	if err != nil {
		panic(middleware.NewNotFound(err.Error()))
	}
	return model

}

func (s *ServiceImpl) FindAllToDo(ctx context.Context) []model.Todo {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	todo := s.Rep.FindAllToDo(ctx, tx)

	return todo
}

func (s *ServiceImpl) CreateActivity(ctx context.Context, request model.Activity) model.Activity {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitorRollback(tx)
	category := model.Activity{
		Title: request.Title,
		Email: request.Email, CreatedAt: request.CreatedAt,
	}
	category = s.Rep.CreateActivity(ctx, tx, category)
	if err != nil {

		panic(err.Error())
	}

	return category

}

func (s *ServiceImpl) UpdateActivity(ctx context.Context, request model.Activity) model.Activity {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer helper.CommitorRollback(tx)

	category, err := s.Rep.FindActivityById(ctx, tx, request.Id)

	if err != nil {
		panic(middleware.NewNotFound(err.Error()))
	}

	category.Title = request.Title
	category = s.Rep.UpdateActivity(ctx, tx, category)
	return category

}

func (s *ServiceImpl) DeleteActivity(ctx context.Context, id int) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	req, err := s.Rep.FindActivityById(ctx, tx, id)
	if err != nil {
		panic(middleware.NewNotFound(err.Error()))
	}

	s.Rep.DeleteActivity(ctx, tx, req.Id)

}

func (s *ServiceImpl) FindAllActivity(ctx context.Context) []model.Activity {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	activities := s.Rep.FindAllActivity(ctx, tx)

	return activities
}

func (s *ServiceImpl) FindActivityById(ctx context.Context, id int) model.Activity {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.FindActivityById(ctx, tx, id)
	if err != nil {
		panic(middleware.NewNotFound(err.Error()))
	}
	return model
}
