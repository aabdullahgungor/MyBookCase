package models

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type CategoryModel struct {
}

func (categoryModel CategoryModel) GetAll() ([]entities.Category, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var categories []entities.Category
		db.Find(&categories)
		return categories, nil
	}
}

func (categoryModel CategoryModel) GetById(id int) (entities.Category, error) {
	db, err := database.GetDB()
	if err != nil {
		return entities.Category{}, err
	} else {
		var category entities.Category
		db.Where("id = ?", id).First(&category)
		return category, nil
	}
}

func (categoryModel CategoryModel) Create(category *entities.Category) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&category)
		return nil
	}
}

func (categoryModel CategoryModel) Edit(category *entities.Category) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&category)
		return nil
	}
}