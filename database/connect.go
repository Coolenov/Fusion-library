package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//var DB *sql.DB

func DbConnect(dbUrl string) *sql.DB {
	var err error
	//dsn := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database")
	return db

}
