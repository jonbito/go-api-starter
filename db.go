package main

import (
	"github.com/jinzhu/gorm"
)

// InitializeDB opens the database connection and stores the value in DB
func InitializeDB() *gorm.DB {
	db, err := gorm.Open(Config.DatabaseDialect, Config.DatabaseURL)
	if err != nil {
		panic(err)
	}
	return db
}
