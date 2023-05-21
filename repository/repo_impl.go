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
)

type RepoImpl struct{}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (r *RepoImpl) CreateToDo(ctx context.Context, tx *sql.Tx, category model.Todo) model.Todo {
	SQL := "INSERT INTO todos(activity_id,title,isactive,created_at,updated_at) VALUES(?,?,?,?,?)"
	rows, err := tx.ExecContext(ctx, SQL, category.ActivityId, category.Title,
		category.Isactive, category.CreatedAt, category.UpdatedAt)
	helper.PanicIfErr(err)
	id, err := rows.LastInsertId()
	if err != nil {
		log.Fatal(err)

	}
	category.Id = int(id)
	return category
}

func (r *RepoImpl) UpdateToDo(ctx context.Context, tx *sql.Tx, category model.Todo) model.Todo {
	SQL := "UPDATE todo SET title=?,updated_at=? WHERE todo_id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Title, category.UpdatedAt, category.Id)
	helper.PanicIfErr(err)

	return category

}

func (r *RepoImpl) DeleteToDo(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM todos WHERE todo_id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAllToDo(ctx context.Context, tx *sql.Tx) []model.Todo {
	sql := "SELECT *FROM todos"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfErr(err)
	defer rows.Close()

	var categoryTodo []model.Todo

	for rows.Next() {
		Todo := model.Todo{}
		current := fmt.Sprint(Todo.CreatedAt)
		current2 := fmt.Sprint(Todo.UpdatedAt)
		err := rows.Scan(&Todo.Id, &Todo.ActivityId, &Todo.Title, &Todo.Priority,
			&Todo.Isactive, &current, &current2)
		Todo.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", current)
		Todo.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", current2)
		helper.PanicIfErr(err)
		categoryTodo = append(categoryTodo, Todo)
	}
	return categoryTodo
}

func (r *RepoImpl) FindTodo(ctx context.Context, tx *sql.Tx, id int) (model.Todo, error) {
	SQL := "SELECT *FROM todos WHERE todo_id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Todo{}
	if rows.Next() {
		current := fmt.Sprint(model.CreatedAt)
		current2 := fmt.Sprint(model.UpdatedAt)
		rows.Scan(&model.Id, &model.ActivityId,
			&model.Title, &model.Priority, &model.Isactive, &current, &current2)
		model.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", current)
		model.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", current2)
		return model, nil
	} else {
		result := fmt.Sprintf("Todo with ID %v Not Found", id)
		return model, errors.New(result)
	}

}

func (m *RepoImpl) CreateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity {
	SQL := `INSERT INTO activities(title,email,created_at,updated_at) VALUES(?,?,?,?)`
	rows, _ := tx.ExecContext(ctx, SQL, category.Title, category.Email, category.CreatedAt, category.UpdatedAt)
	id, err := rows.LastInsertId()
	if err != nil {
		log.Fatal(err)

	}
	category.Id = int(id)
	return category
}

func (r *RepoImpl) UpdateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity {
	SQL := `UPDATE activities set title=?,updated_at=? WHERE activity_id=?`
	_, err := tx.ExecContext(ctx, SQL, category.Title, category.UpdatedAt, category.Id)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) DeleteActivity(ctx context.Context, tx *sql.Tx, id int) {
	SQL := `DELETE FROM activities WHERE activity_id=?`
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAllActivity(ctx context.Context, tx *sql.Tx) []model.Activity {
	SQL := "SELECT *FROM activities"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()
	var activities []model.Activity
	for rows.Next() {
		activity := model.Activity{}
		current := fmt.Sprint(activity.CreatedAt)
		current2 := fmt.Sprint(activity.UpdatedAt)
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Email, &current, &current2)
		log.Printf("current :%v", current)
		activity.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", current)
		activity.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", current2)
		log.Printf("created :%v", activity.CreatedAt)
		if err != nil {
			log.Print(err)
		}
		activities = append(activities, activity)
	}

	return activities
}

func (r *RepoImpl) FindActivityById(ctx context.Context, tx *sql.Tx, id int) (model.Activity, error) {
	SQL := "SELECT *FROM activities WHERE activity_id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Activity{}
	if rows.Next() {
		current := fmt.Sprint(model.CreatedAt)
		current2 := fmt.Sprint(model.UpdatedAt)
		rows.Scan(&model.Id, &model.Title,
			&model.Email, &current, &current2)
		model.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", current)
		model.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", current2)
		return model, nil
	} else {
		result := fmt.Sprintf("Activity with ID %v Not Found", id)
		return model, errors.New(result)
	}

}
