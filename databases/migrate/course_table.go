package migrate

import (
	"os"

	"github.com/Alfeenn/todo-list/model"
)

type CourseTable struct {
	model.Course `gorm:"embedded"`
}

func (CourseTable) TableName() string {
	return os.Getenv("DBNAME") + ".courses"
}
