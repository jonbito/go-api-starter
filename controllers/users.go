package controllers

import (
	"go-api-starter/models"
	"go-api-starter/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserBindingModel struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

// UserController is the controller for the /users route
type UserController struct {
	repo repository.IRepository
}

// NewUserController creates a new controller
func NewUserController(repo repository.IRepository) *UserController {
	c := new(UserController)
	c.repo = repo
	return c
}

// Create creates a new user
func (c UserController) Create(bindingModel interface{}) ControllerResult {

	model, ok := bindingModel.(*UserBindingModel)
	if !ok {
		panic("Cannot convert bindingModel to UserBindingModel")
	}

	// check password length
	if len(model.Password) < 6 {
		return ControllerResult{
			Success:      false,
			ErrorMessage: "Password must be at least 6 characters long",
			Code:         400,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return ControllerResult{
			Success: false,
			Error:   err,
			Code:    500,
		}
	}

	user := models.User{
		Email:    model.Email,
		Password: string(hashedPassword),
	}

	if err = c.repo.Create(&user); err != nil {
		return ControllerResult{
			Success: false,
			Error:   err,
			Code:    500,
		}
	}

	return ControllerResult{
		Success: true,
		Data:    user,
		Code:    201,
	}
}
