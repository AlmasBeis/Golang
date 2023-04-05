package service

import (
	"fmt"
	"thirdassignment/model"
	"thirdassignment/repository"
)

type BookLibraryService struct {
	repo repository.BookLibrary
}

func NewBookLibraryService(repo repository.BookLibrary) *BookLibraryService {
	return &BookLibraryService{repo: repo}
}

// Service >>

func (r *BookLibraryService) AddBook(book model.Book) error {
	if err := r.repo.Create(book); err != nil {
		return err
	}

	return nil
}

func (r *BookLibraryService) GetAllBook() ([]model.Book, error) {
	books, err := r.repo.GetAll()
	if err != nil {
		return []model.Book{}, err
	}

	return books, nil
}

func (r *BookLibraryService) GetBookById(id int) (model.Book, error) {
	book, err := r.repo.GetById(id)
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (r *BookLibraryService) GetBooksByTitle(title string) ([]model.Book, error) {
	books, err := r.repo.GetAllByTitle(title)
	fmt.Println(title)
	if err != nil {
		return []model.Book{}, err
	}

	return books, nil
}

func (r *BookLibraryService) GetOrderedBooksByCost(sort string) ([]model.Book, error) {
	books, err := r.repo.GetAllByCost(sort)
	fmt.Println(sort)
	if err != nil {
		return []model.Book{}, err
	}

	return books, nil
}

func (r *BookLibraryService) SetBook(book model.Book, id int) error {
	if err := r.repo.Update(book, id); err != nil {
		return err
	}
	return nil
}

func (r *BookLibraryService) RemoveBook(id int) error {
	if err := r.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

// << Service
