package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Alfeenn/todo-list/helper"
	"github.com/Alfeenn/todo-list/model"
	"github.com/google/uuid"
)

type RepoImpl struct{}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (r *RepoImpl) CreateToDo(ctx context.Context, tx *sql.Tx, category model.Course) model.Course {
	SQL := "INSERT INTO courses(id,name,price,category,thumbnail) VALUES(?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL,
		category.Id, category.Name, category.Price,
		category.Category, category.Thumbnail)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) UpdateToDo(ctx context.Context, tx *sql.Tx, category model.Course) model.Course {
	SQL := "UPDATE courses SET name=?,price=?,category=?,thumbnail=? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Price, category.Category, category.Thumbnail, category.Id)
	helper.PanicIfErr(err)

	return category

}

func (r *RepoImpl) DeleteToDo(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM user WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) DeleteCourse(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM course WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAllToDo(ctx context.Context, tx *sql.Tx) []model.Course {
	sql := "SELECT *FROM courses"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfErr(err)
	defer rows.Close()

	var categoryCourse []model.Course

	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(&course.Id, &course.Name, &course.Price, &course.Category,
			&course.Thumbnail)
		helper.PanicIfErr(err)
		categoryCourse = append(categoryCourse, course)
	}
	return categoryCourse
}

func (r *RepoImpl) FindTodo(ctx context.Context, tx *sql.Tx, id string) (model.Course, error) {
	SQL := "SELECT *FROM courses WHERE id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Course{}
	if rows.Next() {
		rows.Scan(&model.Id, &model.Name,
			&model.Category, &model.Thumbnail, &model.Price)

		return model, nil
	} else {
		return model, errors.New("no data")
	}

}

func (r *RepoImpl) FindCourseByCategory(ctx context.Context, tx *sql.Tx, category string) (model.Course, error) {
	SQL := "SELECT *FROM courses WHERE category =?"

	rows, err := tx.QueryContext(ctx, SQL, category)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Course{}
	if rows.Next() {
		rows.Scan(&model.Id, &model.Name,
			&model.Price, &model.Category, &model.Thumbnail)

		return model, nil
	} else {
		return model, errors.New("no data")
	}

}

func (m *RepoImpl) Login(ctx context.Context, tx *sql.Tx, category model.User) (model.User, error) {
	SQL := `SELECT id,username,password FROM users WHERE username=?`
	rows, err := tx.QueryContext(ctx, SQL, category.Username)
	helper.PanicIfErr(err)
	defer rows.Close()
	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {

		return user, errors.New("no data")
	}
}

func (r *RepoImpl) Register(ctx context.Context, tx *sql.Tx, category model.User) model.User {
	SQL := "INSERT INTO users(id,username,password,name,age,phone,role) VALUES(?,?,?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL,
		category.Id, category.Username, category.Password,
		category.Name, category.Age, category.Phone, category.Role)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) GetCourse(ctx context.Context, tx *sql.Tx, category model.Class, id string) model.Class {
	SQL := "INSERT INTO class (user_id,course_id) VALUES(?,?)"
	_, err := tx.ExecContext(ctx, SQL,
		category.UserId, id)
	helper.PanicIfErr(err)
	return category
}
