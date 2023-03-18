package models

import (
	"time"
	"fmt"

)

type Book struct {
	ID            uint `gorm:"primary_key, AUTO_INCREMENT"`
	Name          string `gorm:"column:name"`
	Description   string `gorm:"column:description"`
	PublishedDate time.Time `gorm:"column:publishedDate"`
	Edition int `gorm:"column:edition"`
	TotalPages int `gorm:"column:totalPages"`
	Language string `gorm:"column:language"`
	isbn string `gorm:"column:isbn`
	imageUrl string `gorm:"column:imageUrl"`
	AuthorID int `gorm:"column:author_id"`
	Author Author
	ReaderID int `gorm:"column:reader_id"`
	Reader Reader
	Categories []Category `gorm:"many2many:book_has_category"`
	Publishers []Publisher `gorm:"many2many:book_has_publisher"`
}

func (book *Book) TableName() string {
	return "book"
}

func (book Book) ToString() string {
	return fmt.Sprintf("id: %d\nName: %s\nDescription: %s\nPublishedDate: %s\nEdition: %d\nTotalPages: %d\nLanguage: %s\nISBN: %s\nimageUrl: %s", book.ID, book.Name, book.Description, book.PublishedDate, book.Edition, book.TotalPages, book.Language, book.isbn, book.imageUrl)
}