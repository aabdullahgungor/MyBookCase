package models

import "fmt"

type Author struct {
	ID      uint `gorm:"primary_key, AUTO_INCREMENT"`
	Name    string `gorm:"column:name"`
	Books []Book `gorm:"ForeignKey:AuthorID"`
}

func (author *Author) TableName() string {
	return "author"
}

func (author Author) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", author.ID, author.Name)
}
