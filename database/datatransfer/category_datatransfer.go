package datatransfer

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type CategoryData struct {
}

func (categoryData CategoryData) GetAll() ([]entities.Category, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var categories []entities.Category
		db.Find(&categories)
		return categories, nil
	}
}

func (categoryData CategoryData) GetById(id int) ([]entities.Category, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var category []entities.Category
		db.Where("Id = ?", id).Find(&category)
		return category, nil
	}
}

func (categoryData CategoryData) Create(category *entities.Category) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&category)
		return nil
	}
}

func (categoryData CategoryData) Edit(category *entities.Category) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&category)
		return nil
	}
}

func (categoryData CategoryData) Delete(category *entities.Category) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(&category)
		return nil
	}
}

