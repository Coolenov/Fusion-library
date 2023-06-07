package gormDb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DBGORM *gorm.DB

func DbConnect() {
	var err error
	dsn := os.Getenv("DB_URL")
	DBGORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	fmt.Println("Connected to database")
}

//func DbConnect(dbUrl string) *gorm.DB {
//	var err error
//	//dsn := os.Getenv("DB_URL")
//	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Failed to connect to database")
//	}
//	fmt.Println("Connected to database")
//	return db
//}
