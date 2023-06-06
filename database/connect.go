package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func DBConnect(dbUrl string) {
	var err error
	//dsn := os.Getenv("DB_URL")
	DB, err = sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database")
}
