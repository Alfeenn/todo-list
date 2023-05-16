package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

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

func (m *RepoImpl) CreateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity {
	SQL := `INSERT INTO activity(title,email,created_at) VALUES(?,?,?)`

	rows, err := tx.ExecContext(ctx, SQL, category.Title, category.Email, category.CreatedAt)
	helper.PanicIfErr(err)
	id, err := rows.LastInsertId()
	if err != nil {
		log.Fatal(err)

	}
	category.Id = int(id)
	return category
}

func (r *RepoImpl) UpdateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity {
	SQL := `UPDATE activity set title=? WHERE id=?`
	_, err := tx.ExecContext(ctx, SQL, category.Title, category.Id)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) DeleteActivity(ctx context.Context, tx *sql.Tx, id int) {
	SQL := `DELETE FROM activity WHERE id=?`
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAllActivity(ctx context.Context, tx *sql.Tx) []model.Activity {
	SQL := "SELECT *FROM activity"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()
	var activities []model.Activity
	for rows.Next() {
		activity := model.Activity{}
		current := fmt.Sprint(activity.CreatedAt)
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Email, &current)
		log.Printf("current :%v", current)
		activity.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", current)
		log.Printf("created :%v", activity.CreatedAt)
		if err != nil {
			log.Print(err)
		}
		activities = append(activities, activity)
	}

	return activities
}

func (r *RepoImpl) FindActivityById(ctx context.Context, tx *sql.Tx, id int) (model.Activity, error) {
	SQL := "SELECT *FROM activity WHERE id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Activity{}
	if rows.Next() {
		rows.Scan(&model.Id, &model.Title,
			&model.Email, &model.CreatedAt)

		return model, nil
	} else {
		return model, errors.New("no data")
	}

}
