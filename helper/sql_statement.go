package helper

import (
	"fmt"
	"log"

	"github.com/Alfeenn/todo-list/model"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

var dbConfig = model.DBConfig{}

func SQLStatement() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

}

func NewDB() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}
	if err := env.Parse(&dbConfig); err != nil {
		log.Fatal("Unable to parse variables")
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
}
