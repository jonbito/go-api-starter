package controllers

import (
	"go-api-starter/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserBindingModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

// UserController is the controller for the /users route
type UserController struct {
	DB *gorm.DB
}

// NewUserController creates a new controller
func NewUserController(db *gorm.DB) *UserController {
	c := new(UserController)
	c.DB = db
	return c
}

// Create creates a new user
func (c UserController) Create(context *gin.Context) {
	var bindingModel UserBindingModel
	if err := context.ShouldBindJSON(&bindingModel); err != nil {
		context.AbortWithError(400, err)
		return
	}

	// check password length
	if len(bindingModel.Password) < 6 {
		context.AbortWithStatusJSON(400, gin.H{
			"error": "Password must be at least 6 characters long",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bindingModel.Password), bcrypt.DefaultCost)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}

	user := models.User{
		Email:    bindingModel.Email,
		Password: string(hashedPassword),
	}

	if err = c.DB.Create(&user).Error; err != nil {
		context.AbortWithError(500, err)
		return
	}

	context.JSON(201, gin.H{
		"data": user,
	})
}
