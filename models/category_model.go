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
		db.Preload("Books").Find(&categories)
		return categories, nil
	}
}