package migrate

import (
	"os"

	"github.com/Alfeenn/todo-list/model"
)

type ActivityTable struct {
	model.Activity `gorm:"embedded"`
}

func (ActivityTable) TableName() string {
	return os.Getenv("MYSQL_DBNAME") + ".activity"
}
