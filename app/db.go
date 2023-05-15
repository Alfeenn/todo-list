package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Alfeenn/online-learning/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	CreateDB()
	statement := helper.SQLStatement()
	db, err := sql.Open("mysql", statement)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	return db
}

func CreateDB() {
	initDB := helper.NewDB()
	dbName := os.Getenv("DBNAME")
	db, err := sql.Open("mysql", initDB)
	if err != nil {
		log.Fatal("connection error")
	}
	defer db.Close()
	createDBCommand := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	_, err = db.Exec(createDBCommand)
	if err != nil {
		log.Fatal(err)
	}
}
