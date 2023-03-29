package entities

import (
	"fmt"
)

type Publisher struct {
	ID            uint `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	PublisherName string `gorm:"column:publisher" json:"publisher_name"`
	Books []Book `gorm:"many2many:book_has_publisher"`
}

func (publisher *Publisher) TableName() string {
	return "publisher"
}

func (publisher Publisher) ToString() string {
	return fmt.Sprintf("id: %d\nPublisher: %s", publisher.ID, publisher.PublisherName)
}