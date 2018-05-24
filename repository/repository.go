package repository

import (
	"github.com/jinzhu/gorm"
)

// IRepository is the interface for all repositories
type IRepository interface {
	Create(data interface{}) error
	Find(out interface{}, where ...interface{}) error
	First(out interface{}, where ...interface{}) error
}

// GormRepository is a repository that uses GORM and upholds the IRepository interface
type GormRepository struct {
	DB *gorm.DB
}

// Create creates a new database entry of data
func (r *GormRepository) Create(data interface{}) error {
	return r.DB.Create(data).Error
}

// Find retrieves a database entry and places it in data
func (r *GormRepository) Find(out interface{}, where ...interface{}) error {
	return r.DB.Find(out, where...).Error
}

func (r *GormRepository) First(out interface{}, where ...interface{}) error {
	return r.DB.First(out, where...).Error
}

// NewGormRepository returns a new GormRepository
func NewGormRepository(db *gorm.DB) *GormRepository {
	r := new(GormRepository)
	r.DB = db
	return r
}
