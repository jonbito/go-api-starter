package controllers

import (
	"go-api-starter/models"
	"go-api-starter/repository"
	"strings"
	"testing"

	"gopkg.in/go-playground/validator.v8"
)

type MockUserRepository struct {
	repository.MockRepository
}

func (r *MockUserRepository) Create(data interface{}) error {
	user, ok := data.(*models.User)
	if !ok {
		panic("Could not convert data to user")
	}
	user.ID = 1
	user.Email = "asdf@asdf.com"
	user.FirstName = "Joe"
	return nil
}

func TestCreate(t *testing.T) {
	c := NewUserController(new(MockUserRepository))
	bindingModel := UserBindingModel{
		Email:    "asdf@asdf",
		Password: "asdf",
	}
	result := c.Create(&bindingModel)
	if result.ErrorMessage != "Password must be at least 6 characters long" {
		t.Error("Password is less than 6 characters")
	}

	bindingModel = UserBindingModel{
		Email:    "asdf@asdf.com",
		Password: "asdfasdf",
	}
	result = c.Create(&bindingModel)
	user, ok := result.Data.(models.User)
	if !ok {
		panic("Could not convert result to user")
	}
	if user.FirstName != "Joe" {
		t.Error("Did not receive the correct user")
	}
}

func TestUserBindingModel(t *testing.T) {
	bindingModel := UserBindingModel{
		Email:    "asdf",
		Password: "asdf",
	}
	v := validator.New(&validator.Config{
		TagName: "binding",
	})

	err := v.Struct(bindingModel)
	if err == nil {
		t.Error("No error")
	}
	if err != nil && !strings.Contains(err.Error(), "UserBindingModel.Email") {
		t.Error("Email validation failed")
	}

	bindingModel = UserBindingModel{
		Email:    "asdf@asdf.com",
		Password: "asdf",
	}
	err = v.Struct(bindingModel)
	if err != nil {
		t.Error("Validation should have passed")
	}
}
