package users

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindOne(username string) (User, error)
	CreateUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOne(Username string) (User, error) {
	var user User
	err := r.db.Where("username = ?", Username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// handle record not found
		err = gorm.ErrRecordNotFound
	}

	return user, err
}

func (r *repository) CreateUser(user User) (User, error) {

	err := r.db.Create(&user).Error

	return user, err
}
