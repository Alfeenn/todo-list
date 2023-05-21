package migrate

import (
	"os"

	"github.com/Alfeenn/todo-list/model"
)

type TodoTable struct {
	model.Todos `gorm:"embedded"`
}

func (TodoTable) TableName() string {
	return os.Getenv("MYSQL_DBNAME") + ".todos"
}
