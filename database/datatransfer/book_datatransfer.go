package datatransfer

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type BookData struct {
}

func (bookData BookData) GetAll() ([]entities.Book, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var books []entities.Book
		db.Find(&books)
		return books, nil
	}
}

func (bookData BookData) GetById(id int) ([]entities.Book, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var book []entities.Book
		db.Where("Id = ?", id).Find(&book)
		return book, nil
	}
}

func (bookData BookData) Create(book *entities.Book) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&book)
		return nil
	}
}

func (bookData BookData) Edit(book *entities.Book) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&book)
		return nil
	}
}

func (bookData BookData) Delete(book *entities.Book) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(&book)
		return nil
	}
}


