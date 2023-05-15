package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/Alfeenn/todo-list/helper"
	"github.com/Alfeenn/todo-list/model"
	"github.com/Alfeenn/todo-list/model/web"
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

func (s *ServiceImpl) DeleteUser(ctx context.Context, id string) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	s.Rep.DeleteToDo(ctx, tx, id)
}

func (s *ServiceImpl) DeleteToDo(ctx context.Context, id string) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	req, err := s.Rep.FindTodo(ctx, tx, id)
	s.Rep.DeleteCourse(ctx, tx, req.Id)
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

func (s *ServiceImpl) FindCourseByCategory(ctx context.Context, id string) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.FindCourseByCategory(ctx, tx, id)
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

func (s *ServiceImpl) Login(ctx context.Context, request web.RequestLogin) web.CatResp {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitorRollback(tx)
	category := model.User{
		Username: request.Username,
		Password: request.Password,
	}
	category, err = s.Rep.Login(ctx, tx, category)
	if err != nil {

		panic(err.Error())
	}

	return helper.ConvertModel(category)

}

func (s *ServiceImpl) Register(ctx context.Context, request web.CategoryRequest) web.CatResp {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer helper.CommitorRollback(tx)
	category := model.User{
		Id:       request.Id,
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
		Age:      request.Age,
		Phone:    request.Phone,
		Role:     request.Role,
	}
	if category.Role == "" {
		category.Role = "user"
	}
	category = s.Rep.Register(ctx, tx, category)
	return helper.ConvertModel(category)
}

func (s *ServiceImpl) GetCourse(ctx context.Context, req model.Class, id string) model.Class {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	Class := s.Rep.GetCourse(ctx, tx, req, id)
	Class.CourseId = id
	return Class

}
