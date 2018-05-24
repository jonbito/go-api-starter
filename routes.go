package main

import (
	"go-api-starter/controllers"
	"go-api-starter/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Routes creates all routes used in the API
func Routes(r *gin.Engine, db *gorm.DB) {
	// we use authMiddleware everywhere so we instantiate it here
	authMiddleware := middleware.CreateAuthMiddleware(db, Config.JWTRealm, Config.JWTSecret, Config.JWTTimeout, Config.JWTMaxRefresh)

	// Do not require authentication for these routes
	userController := controllers.NewUserController(db)
	r.POST("/auth/login", authMiddleware.LoginHandler, middleware.RateLimiter(Config.RateLimiterLoginPeriod, Config.RateLimiterLoginLimit))
	r.POST("/users/create", userController.Create)

	// Require JWT Authentication
	auth := r.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/auth/refresh", authMiddleware.RefreshHandler)
	}
}
