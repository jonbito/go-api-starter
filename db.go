package main

import (
	"github.com/jinzhu/gorm"
)

// DB contains the currently open database connection
var DB *gorm.DB

// InitializeDB opens the database connection and stores the value in DB
func InitializeDB() {
	db, err := gorm.Open(Config.DatabaseDialect, Config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	DB = db
}
