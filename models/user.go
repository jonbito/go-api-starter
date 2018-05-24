package models

import (
	"github.com/jinzhu/gorm"
)

// User defines an application user
type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(50);unique_index;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"type:varchar(30)"`
	LastName  string `gorm:"type:varchar(30)"`
	Role      string `gorm:"type:varchar(15)"`
}
