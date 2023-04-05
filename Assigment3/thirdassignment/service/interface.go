package service

import (
	"thirdassignment/model"
	"thirdassignment/repository"
)

type BookLibrary interface {
	AddBook(book model.Book) error
	GetAllBook() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
	GetBooksByTitle(title string) ([]model.Book, error)
	GetOrderedBooksByCost(sort string) ([]model.Book, error)
	SetBook(book model.Book, id int) error
	RemoveBook(id int) error
}

type Service struct {
	BookLibrary
}

func NewService(repo *repository.Repository) *Service {
	return &Service{BookLibrary: NewBookLibraryService(repo.BookLibrary)}
}
