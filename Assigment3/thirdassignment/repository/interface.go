package repository

import (
	"gorm.io/gorm"
	"thirdassignment/model"
)

type BookLibrary interface {
	Create(book model.Book) error
	GetAll() ([]model.Book, error)
	GetById(id int) (model.Book, error)
	GetAllByTitle(title string) ([]model.Book, error)
	GetAllByCost(sort string) ([]model.Book, error)
	Update(book model.Book, id int) error
	Delete(id int) error
}

type Repository struct {
	BookLibrary
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{NewBookLibraryCRUD(db)}
}
