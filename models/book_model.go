package models

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type BookModel struct {
}


func (bookModel BookModel) GetAll() ([]entities.Book, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var books []entities.Book
		db.Find(&books)
		return books, nil
	}
}


func (bookModel BookModel) GetById(id int) (entities.Book, error) {
	db, err := database.GetDB()
	if err != nil {
		return entities.Book{}, err
	} else {
		var book entities.Book
		db.Where("id = ?", id).First(&book)
		return book, nil
	}
}

func (bookModel BookModel) Create(book *entities.Book) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&book)
		return nil
	}
}


func (bookModel BookModel) Edit(book *entities.Book) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&book)
		return nil
	}
}