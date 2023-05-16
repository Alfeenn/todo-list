package migrate

import (
	"os"

	"github.com/Alfeenn/todo-list/model"
)

type UserTable struct {
	model.User `gorm:"embedded"`
}

func (UserTable) TableName() string {
	return os.Getenv("DBNAME") + ".users"
}
