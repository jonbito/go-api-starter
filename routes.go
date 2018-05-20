package main

import (
	"github.com/gin-gonic/gin"
	"go-api-starter/controllers"
	"go-api-starter/middleware"
)

// Routes creates all routes used in the API
func Routes(r *gin.Engine) {
	authMiddleware := middleware.CreateAuthMiddleware(DB)

	r.POST("/auth/login", authMiddleware.LoginHandler)
	userController := controllers.NewUserController(DB)
	r.POST("/users/create", userController.Create)

	auth := r.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/auth/refresh", authMiddleware.RefreshHandler)

	}
}
