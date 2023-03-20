package business

import (
	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/database/datatransfer"
)

type Category struct{
}

func (c Category) GetAll() ([]models.CategoryModel, error) {
	categories, err := datatransfer.CategoryData{}.GetAll()
	if err != nil {
		return &[]entities.Category{}, err
	}

	var newCategory []models.CategoryModel
	var dataCategory []models.CategoryModel

	for _,category := range categories {
		newCategory.ID = category.ID
		newCategory.CategoryName = category.CategoryName 

		dataCategory = append(dataCategory, newCategory)
	}
	return &dataCategory, nil
}

func (c Category) GetById(id int) ([]models.CategoryModel, error) {
	category, err := datatransfer.CategoryData{}.GetById(id)
	if err != nil {
		return &[]entities.Category{}, err
	}

	var newCategory []models.CategoryModel

	newCategory.ID = category.ID
	newCategory.CategoryName = category.CategoryName

	return &newCategory, nil

}

func (c Category) Create(newCategory *models.CategoryModel) error {
	err := datatransfer.CategoryData{}.Create(&newCategory)
	if err != nil {
		return err
	}
	return nil
}

func (c Category) Edit(editedCategory *models.CategoryModel) error {
	err := datatransfer.CategoryData{}.Edit(&editedCategory)
	if err != nil {
		return err
	}
	return nil
}

func (c Category) Delete(deletedCategory *models.CategoryModel) error {
	err := datatransfer.CategoryData{}.Delete(&deletedCategory)
	if err != nil {
		return err
	}
	return nil
}

