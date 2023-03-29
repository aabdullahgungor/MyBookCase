package entities

import "fmt"

type Author struct {
	ID      uint `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Books []Book `gorm:"ForeignKey:AuthorID" json:"author_id"`
}

func (author *Author) TableName() string {
	return "author"
}

func (author Author) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", author.ID, author.Name)
}
