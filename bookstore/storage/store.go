package storage

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	db *gorm.DB
)

func Connect() *gorm.DB {
	user := os.Getenv("USER")
	pwd := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local", user, pwd)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
