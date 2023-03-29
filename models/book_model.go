package models

import (
	"errors"
	"reflect"

	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

var (
	ErrBookIDIsNotValid    = errors.New("id is not valid")
	ErrBookNameIsNotEmpty = errors.New("Book name cannot be empty")
	ErrBookNotFound   = errors.New("the book cannot be found")
	ErrBookAuthorIDIsNotValid = errors.New("Author_id is not valid")
	ErrBookReaderIDIsNotValid = errors.New("Reader_id is not valid")
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
	if id <= 0 || reflect.TypeOf(id).Kind() != reflect.Int{
		return entities.Book{}, ErrBookIDIsNotValid
	}

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

	switch  {
	case book.Name == "":
		return ErrBookNameIsNotEmpty
	case book.AuthorID <= 0 : // || reflect.TypeOf(book.AuthorID).Kind() != reflect.Int
		return ErrBookAuthorIDIsNotValid
	case book.ReaderID <= 0 :
		return ErrBookReaderIDIsNotValid
	default:
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Create(&book)
			return nil
		}

	}
	
}


func (bookModel BookModel) Edit(book *entities.Book) error {

	switch  {
	case book.Name == "":
		return ErrBookNameIsNotEmpty
	case book.AuthorID <= 0 : // || reflect.TypeOf(book.AuthorID).Kind() != reflect.Int
		return ErrBookAuthorIDIsNotValid
	case book.ReaderID <= 0 :
		return ErrBookReaderIDIsNotValid
	default:
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Save(&book)
			return nil
		}
	}
	
}

func (bookModel BookModel) Delete(id int) error {
	book, err := bookModel.GetById(id)
	if err != nil {
		return err
	}
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	db.Delete(book)
	return nil
}