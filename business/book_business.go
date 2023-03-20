package business

import (
	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/database/datatransfer"
)

type Book struct{
}

func (b Book) GetAll() ([]models.BookModel, error) {
	books, err := datatransfer.BookData{}.GetAll()
	if err != nil {
		return &[]entities.Book{}, err
	}

	var newBook []models.BookModel
	var dataBook []models.BookModel

	for _,book := range books {
		
		newBook.ID = book.ID
		newBook.Name = book.Name
		newBook.Description = book.Description
		newBook.PublishedDate = book.PublishedDate
		newBook.Edition = book.Edition
		newBook.TotalPages = book.TotalPages
		newBook.Language = book.Language
		newBook.isbn = book.isbn
		newBook.imageUrl = book.imageUrl
		newBook.AuthorID = book.AuthorID
		newBook.ReaderID = book.ReaderID

		dataBook = append(dataBook, newBook)
	}
	return &dataBook, nil
}

func (b Book) GetById(id int) ([]models.BookModel, error) {
	book, err := datatransfer.BookData{}.GetById(id)
	if err != nil {
		return &[]entities.Book{}, err
	}

	var newBook []models.BookModel

	newBook.ID = book.ID
	newBook.Name = book.Name
	newBook.Description = book.Description
	newBook.PublishedDate = book.PublishedDate
	newBook.Edition = book.Edition
	newBook.TotalPages = book.TotalPages
	newBook.Language = book.Language
	newBook.isbn = book.isbn
	newBook.imageUrl = book.imageUrl
	newBook.AuthorID = book.AuthorID
	newBook.ReaderID = book.ReaderID

	return &newBook, nil

}

func (b Book) Create(newBook *models.BookModel) error {
	err := datatransfer.BookData{}.Create(&newBook)
	if err != nil {
		return err
	}
	return nil
}

func (b Book) Create(editedBook *models.BookModel) error {
	err := datatransfer.BookData{}.Edit(&editedBook)
	if err != nil {
		return err
	}
	return nil
}

func (a Book) Delete(deletedBook *models.BookModel) error {
	err := datatransfer.BookData{}.Delete(&deletedBook)
	if err != nil {
		return err
	}
	return nil
}

