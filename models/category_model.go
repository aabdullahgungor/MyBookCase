package models

import (
	"errors"
	"reflect"

	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

var (
	ErrCategoryIDIsNotValid    = errors.New("id is not valid")
	ErrCategoryNameIsNotEmpty = errors.New("Category name cannot be empty")
	ErrCategoryNotFound   = errors.New("the category cannot be found")
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
	if id <= 0 || reflect.TypeOf(id).Kind() != reflect.Int{
		return entities.Category{}, ErrCategoryIDIsNotValid
	}

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
	
	if category.CategoryName == "" {
		return ErrCategoryNameIsNotEmpty
	}else {

		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Create(&category)
			return nil
		}
	}
	
}

func (categoryModel CategoryModel) Edit(category *entities.Category) error {

	if category.CategoryName == "" {
		return ErrCategoryNameIsNotEmpty
	}else {

		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Save(&category)
			return nil
		}
	}
	
}

func (categoryModel CategoryModel) Delete(id int) error {
	category, err := categoryModel.GetById(id)
	if err != nil {
		return err
	}
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	db.Delete(category)
	return nil
}