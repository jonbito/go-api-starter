package main

import (
	"go-api-starter/models"
)

// Migrate automigrates the database
// Put all your models that you want automigrated here
func Migrate() {
	DB.AutoMigrate(&models.User{})
}
