package entities

import (
	"fmt"
)

type Category struct {
	ID           uint `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	CategoryName string `gorm:"column:category" json:"category_name"`
	Books []Book `gorm:"many2many:book_has_category"`
}

func (category *Category) TableName() string {
	return "category"
}

func (category Category) ToString() string {
	return fmt.Sprintf("id: %d\nCategory: %s", category.ID, category.CategoryName)
}
