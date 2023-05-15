package databases

import (
	"fmt"

	"github.com/Alfeenn/online-learning/app"
	"github.com/Alfeenn/online-learning/databases/migrate"
	"github.com/Alfeenn/online-learning/helper"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Tables() []interface{} {
	return []interface{}{
		&migrate.UserTable{},
		&migrate.CourseTable{},
		&migrate.ClassTable{},
	}
}

func MigrationDB() *gorm.DB {
	err := godotenv.Load(".env")
	helper.PanicIfErr(err)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{ // use existing connection
		Conn: app.NewDB(),
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Migration DB Error: ", err.Error())
	}

	return gormDB
}
