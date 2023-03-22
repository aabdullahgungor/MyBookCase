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
		db.Preload("Categories").Find(&books)
		return books, nil
	}
}


