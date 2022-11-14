package book

import (
	"fmt"

	"gorm.io/gorm"
)

// interface Repository to use all func inside
type Repository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Find all function
func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	if len(books) == 0 {
		err = gorm.ErrRecordNotFound
	}

	return books, err
}

// Find data by id
func (r *repository) FindByID(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Error
	if err != nil {
		fmt.Println(err)
	}

	return book, err
}

// Create data
func (r *repository) Create(book Book) (Book, error) {

	err := r.db.Create(&book).Error

	return book, err
}
