package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/Alfeenn/todo-list/helper"
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

func (s *ServiceImpl) CreateToDo(ctx context.Context, req model.Course) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	request := model.Course{
		Id:        req.Id,
		Name:      req.Name,
		Price:     req.Price,
		Category:  req.Category,
		Thumbnail: req.Thumbnail,
	}
	Course := s.Rep.CreateToDo(ctx, tx, request)

	return Course

}

func (s *ServiceImpl) UpdateToDo(ctx context.Context, req model.Course) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	id := req.Id
	findId, err := s.Rep.FindTodo(ctx, tx, id)
	helper.PanicIfErr(err)
	updateCourse := s.Rep.UpdateToDo(ctx, tx, findId)
	return updateCourse
}

func (s *ServiceImpl) DeleteToDo(ctx context.Context, id string) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	req, err := s.Rep.FindTodo(ctx, tx, id)
	s.Rep.FindTodo(ctx, tx, req.Id)
	helper.PanicIfErr(err)
	s.Rep.DeleteToDo(ctx, tx, req.Id)

}

func (s *ServiceImpl) FindTodo(ctx context.Context, id string) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.FindTodo(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	return model

}

func (s *ServiceImpl) FindAllToDo(ctx context.Context) []model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	slicemodel := s.Rep.FindAllToDo(ctx, tx)

	var sliceCourse []model.Course

	for _, v := range slicemodel {
		sliceCourse = append(sliceCourse, v)
	}
	return sliceCourse
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
		panic(err.Error())
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
	helper.PanicIfErr(err)
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
		log.Print(err.Error())
	}
	return model
}
