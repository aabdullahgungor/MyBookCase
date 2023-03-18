package models
import 
(
	"fmt"

)

type Category struct {
	ID           uint `gorm:"primary_key, AUTO_INCREMENT"`
	CategoryName string `gorm:"column:category"`
	Books []Book `gorm:"many2many:book_has_category"`
}

func (category *Category) TableName() string {
	return "category"
}

func (category Category) ToString() string {
	return fmt.Sprintf("id: %d\nCategory: %s", category.ID, category.CategoryName)
}
