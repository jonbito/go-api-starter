package main

import (
	"github.com/gin-gonic/gin"
	"go-api-starter/middleware"
)

func main() {
	// initialize the db
	InitializeDB()

	// automigrate the database
	Migrate()

	r := gin.Default()
	r.Use(middleware.ErrorMiddleware)

	// setup the routes
	Routes(r)

	r.Run(":" + GetEnv("PORT", "8080"))
}
