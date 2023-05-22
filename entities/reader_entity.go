package entities

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Reader struct {
	ID       uint   `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Books    []Book `gorm:"ForeignKey:ReaderID"`
}

type Authentication struct { // Authentication is for login data.
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct { // Token is for storing token information for correct login credentials.
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func (reader *Reader) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	reader.Password = string(bytes)
	return nil
}
func (reader *Reader) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(reader.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (reader *Reader) TableName() string {
	return "reader"
}

func (reader Reader) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", reader.ID, reader.Name)
}
