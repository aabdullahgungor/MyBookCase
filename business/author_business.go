package business

import (
	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/database/datatransfer"
)

type Author struct{
}

func (a Author) GetAll() ([]models.AuthorModel, error) {
	authors, err := datatransfer.AuthorData{}.GetAll()
	if err != nil {
		return &[]entities.Author{}, err
	}

	var newAuthor []models.AuthorModel
	var dataAuthor []models.AuthorModel

	for _,author := range authors {
		newAuthor.ID = author.ID
		newAuthor.Name = author.Name
		dataAuthor = append(dataAuthor, newAuthor)
	}
	return &dataAuthor, nil
}

func (a Author) GetById(id int) ([]models.AuthorModel, error) {
	author, err := datatransfer.AuthorData{}.GetById(id)
	if err != nil {
		return &[]entities.Author{}, err
	}

	var newAuthor []models.AuthorModel

	newAuthor.ID = author.ID
	newAuthor.Name = author.Name

	return &newAuthor, nil

}

func (a Author) Create(newAuthor *models.AuthorModel) error {
	err := datatransfer.AuthorData{}.Create(&newAuthor)
	if err != nil {
		return err
	}
	return nil
}

func (a Author) Edit(editedAuthor *models.AuthorModel) error {
	err := datatransfer.AuthorData{}.Edit(&editedAuthor)
	if err != nil {
		return err
	}
	return nil
}

func (a Author) Delete(deletedAuthor *models.AuthorModel) error {
	err := datatransfer.AuthorData{}.Delete(&deletedAuthor)
	if err != nil {
		return err
	}
	return nil
}

