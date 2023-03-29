package entities

import (
	"fmt"
	"time"
)

type Book struct {
	ID            uint `gorm:"primary_key, AUTO_INCREMENT" json:"id" `
	Name          string `gorm:"column:name" json:"name"`
	Description   string `gorm:"column:description" json:"description"`
	PublishedDate time.Time `gorm:"column:publishedDate" json:"published_date"`
	Edition int `gorm:"column:edition" json:"edition"`
	TotalPages int `gorm:"column:totalPages" json:"total_pages"`
	Language string `gorm:"column:language" json:"language"`
	Isbn string `gorm:"column:isbn" json:"isbn"`  // Export a field by starting it with an uppercase letter.
	ImageUrl string `gorm:"column:imageUrl" json:"image_url"`
	AuthorID int `gorm:"column:author_id" json:"author_id"`
	Author Author
	ReaderID int `gorm:"column:reader_id" json:"reader_id"`
	Reader Reader
	Categories []Category `gorm:"many2many:book_has_category"`
	Publishers []Publisher `gorm:"many2many:book_has_publisher"`
}

func (book *Book) TableName() string {
	return "book"
}

func (book Book) ToString() string {
	return fmt.Sprintf("id: %d\nName: %s\nDescription: %s\nPublishedDate: %s\nEdition: %d\nTotalPages: %d\nLanguage: %s\nISBN: %s\nimageUrl: %s", book.ID, book.Name, book.Description, book.PublishedDate, book.Edition, book.TotalPages, book.Language, book.Isbn, book.ImageUrl)
}