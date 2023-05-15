package migrate

import (
	"os"

	"github.com/Alfeenn/online-learning/model"
)

type ClassTable struct {
	model.Class `gorm:"embedded"`
}

func (ClassTable) TableName() string {
	return os.Getenv("DBNAME") + ".class"
}
