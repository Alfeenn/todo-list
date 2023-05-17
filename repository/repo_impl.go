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
	SQL := "INSERT INTO todo(activity_id,title,isactive) VALUES(?,?,?)"
	rows, err := tx.ExecContext(ctx, SQL, category.ActivityId, category.Title,
		category.Isactive)
	id, err := rows.LastInsertId()
	if err != nil {
		log.Fatal(err)

	}
	category.Id = int(id)
	return category
}

func (r *RepoImpl) UpdateToDo(ctx context.Context, tx *sql.Tx, category model.Todo) model.Todo {
	SQL := "UPDATE todo SET title=? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Title, category.Id)
	helper.PanicIfErr(err)

	return category

}

func (r *RepoImpl) DeleteToDo(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM todo WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAllToDo(ctx context.Context, tx *sql.Tx) []model.Todo {
	sql := "SELECT *FROM todo"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfErr(err)
	defer rows.Close()

	var categoryTodo []model.Todo

	for rows.Next() {
		Todo := model.Todo{}
		err := rows.Scan(&Todo.Id, &Todo.ActivityId, &Todo.Title, &Todo.Priority,
			&Todo.Isactive)
		helper.PanicIfErr(err)
		categoryTodo = append(categoryTodo, Todo)
	}
	return categoryTodo
}

func (r *RepoImpl) FindTodo(ctx context.Context, tx *sql.Tx, id int) (model.Todo, error) {
	SQL := "SELECT *FROM todo WHERE id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Todo{}
	if rows.Next() {
		rows.Scan(&model.Id, &model.ActivityId,
			&model.Title, &model.Priority, &model.Isactive)

		return model, nil
	} else {
		return model, errors.New("no data")
	}

}

func (m *RepoImpl) CreateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity {
	SQL := `INSERT INTO activities(title,email,created_at) VALUES(?,?,?)`
	category.CreatedAt.Add(time.Since(time.Now()))
	rows, _ := tx.ExecContext(ctx, SQL, category.Title, category.Email, category.CreatedAt)
	id, err := rows.LastInsertId()
	if err != nil {
		log.Fatal(err)

	}
	category.Id = int(id)
	return category
}

func (r *RepoImpl) UpdateActivity(ctx context.Context, tx *sql.Tx, category model.Activity) model.Activity {
	SQL := `UPDATE activities set title=? WHERE id=?`
	_, err := tx.ExecContext(ctx, SQL, category.Title, category.Id)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) DeleteActivity(ctx context.Context, tx *sql.Tx, id int) {
	SQL := `DELETE FROM activities WHERE id=?`
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
	SQL := "SELECT *FROM activities WHERE id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Activity{}
	if rows.Next() {
		rows.Scan(&model.Id, &model.Title,
			&model.Email, &model.CreatedAt)

		return model, nil
	} else {
		result := fmt.Sprintf("Activity with ID %v Not Found", id)
		return model, errors.New(result)
	}

}
