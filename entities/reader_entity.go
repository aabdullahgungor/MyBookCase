package entities

import "fmt"

type Reader struct {
	ID       uint   `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Books    []Book `gorm:"ForeignKey:ReaderID"`
}

func (reader *Reader) TableName() string {
	return "reader"
}

func (reader Reader) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", reader.ID, reader.Name)
}
