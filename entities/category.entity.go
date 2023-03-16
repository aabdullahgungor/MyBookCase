package models

type Category struct {
	ID           int `gorm:"primary_key, AUTO_INCREMENT"`
	CategoryName string `gorm:"column:Category"`
	Books []Book `gorm:"many2many:book_category"`
}