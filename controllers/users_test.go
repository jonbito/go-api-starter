package controllers

import (
	"go-api-starter/models"
	"go-api-starter/repository"
	"testing"
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
}
