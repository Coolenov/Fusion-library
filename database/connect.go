package database

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func DBConnect() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database")
}
