package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-api-starter/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
		context.AbortWithStatusJSON(400, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bindingModel.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user := models.User{
		Email:    bindingModel.Email,
		Password: string(hashedPassword),
	}

	if err = c.DB.Create(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
