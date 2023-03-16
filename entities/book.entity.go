package models

import (
	"time"
)

type Book struct {
	ID            uint `gorm:"primary_key, AUTO_INCREMENT"`
	Name          string `gorm:"column:Name"`
	Description   string `gorm:"column:Description"`
	PublishedDate time.Time `gorm:"column:PublishedDate"`
	Edition int `gorm:"column:Edition"`
	TotalPages int `gorm:"column:TotalPages"`
	Language string `gorm:"column:Language"`
	isbn string `gorm:"column:ISBN"`
	imageUrl string `gorm:"column:ImageURL"`
	Categories []Category `gorm:"many2many:book_category"`
	Publishers []Publisher `gorm:"many2many:book_publisher"`
	AuthorID int `gorm:"column:Author_id"`
	Author Author
	ReaderID int `gorm:"column:Reader_id"`
	Reader Reader
}
