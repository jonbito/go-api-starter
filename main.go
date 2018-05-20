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

	// setup gin
	r := gin.Default()
	r.Use(middleware.ErrorMiddleware)

	// Add rate limit middleware
	r.Use(middleware.RateLimiter(Config.RateLimiterPeriod, Config.RateLimiterLimit))

	// setup the routes
	Routes(r)

	// use the environment variable PORT or 8080 if PORT is not defined
	r.Run(":" + GetEnv("PORT", "8080"))
}
