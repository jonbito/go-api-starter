package main

import (
	"github.com/gin-gonic/gin"
	"go-api-starter/controllers"
	"go-api-starter/middleware"
)

// Routes creates all routes used in the API
func Routes(r *gin.Engine) {
	// we use authMiddleware everywhere so we instantiate it here
	authMiddleware := middleware.CreateAuthMiddleware(DB, Config.JWTRealm, Config.JWTSecret, Config.JWTTimeout, Config.JWTMaxRefresh)

	// Do not require authentication for these routes
	r.POST("/auth/login", authMiddleware.LoginHandler)
	userController := controllers.NewUserController(DB)
	r.POST("/users/create", userController.Create)

	// Require JWT Authentication
	auth := r.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/auth/refresh", authMiddleware.RefreshHandler)

	}
}
