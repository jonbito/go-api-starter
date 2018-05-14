package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB contains the currently open database connection
var DB *gorm.DB

// InitializeDB opens the database connection and stores the value in DB
func InitializeDB() {
	db, err := gorm.Open("postgres", GetEnv("DATABASE_URL", "postgres://postgres:admin@localhost:5432/test?sslmode=disable"))
	if err != nil {
		panic(err)
	}
	DB = db
}
