package models

type Reader struct {
	ID      int `gorm:"primary_key, AUTO_INCREMENT"`
	Name    string `gorm:"column:Name"`
	Surname string `gorm:"column:Surname"`
	Books []Book `gorm:"ForeignKey:ReaderID"`
}