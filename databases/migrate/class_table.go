package migrate

import (
	"os"

	"github.com/Alfeenn/todo-list/model"
)

type TodoTable struct {
	model.Todo `gorm:"embedded"`
}

func (TodoTable) TableName() string {
	return os.Getenv("MYSQL_DBNAME") + ".todo"
}
