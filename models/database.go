package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=musicstore password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Track{})

	return db
}
