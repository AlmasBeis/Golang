package repository

import (
	"gorm.io/gorm"
	"thirdassignment/model"
)

type BookLibraryCRUD struct {
	db *gorm.DB
}

func NewBookLibraryCRUD(db *gorm.DB) *BookLibraryCRUD {
	return &BookLibraryCRUD{db: db}
}

// CRUD >>
// Create

func (d *BookLibraryCRUD) Create(book model.Book) error {
	if result := d.db.Create(&book); result.Error != nil {
		return result.Error
	}

	return nil
}

// Read

func (d *BookLibraryCRUD) GetAll() ([]model.Book, error) {
	var books []model.Book

	if result := d.db.Find(&books); result.Error != nil {
		return []model.Book{}, result.Error
	}

	return books, nil
}

func (d *BookLibraryCRUD) GetById(id int) (model.Book, error) {
	var book model.Book

	if result := d.db.First(&book, id); result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

func (d *BookLibraryCRUD) GetAllByTitle(title string) ([]model.Book, error) {
	var books []model.Book

	// title IN ?, []string{title} this one should also work, I have not tried yet
	if result := d.db.Where("title LIKE ?", "%"+title+"%").Find(&books); result.Error != nil {
		return []model.Book{}, result.Error
	}

	return books, nil
}

func (d *BookLibraryCRUD) GetAllByCost(sort string) ([]model.Book, error) {
	var books []model.Book
	if result := d.db.Order("cost " + sort).Find(&books); result.Error != nil {
		return []model.Book{}, result.Error
	}

	return books, nil
}

// Update

func (d *BookLibraryCRUD) Update(book model.Book, id int) error {
	var newBook model.Book

	if result := d.db.First(&newBook, id); result.Error != nil {
		return result.Error
	}

	// have not tried this one newBook = book, because of book's id
	newBook.Cost = book.Cost
	newBook.Title = book.Title
	newBook.Description = book.Description

	if result := d.db.Save(&newBook); result.Error != nil {
		return result.Error
	}

	return nil
}

// Delete

func (d *BookLibraryCRUD) Delete(id int) error {
	if result := d.db.Delete(&model.Book{}, id); result.Error != nil {
		return result.Error
	}

	return nil
}

// << CRUD
