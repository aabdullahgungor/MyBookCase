package models

type Publisher struct {
	ID            int `gorm:"primary_key, AUTO_INCREMENT"`
	PublisherName string `gorm:"column:Publisher"`
	Books []Book `gorm:"many2many:book_publisher"`
}