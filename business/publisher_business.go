package business

import (
	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/database/datatransfer"
)

type Publisher struct{
}

func (p Publisher) GetAll() ([]models.PublisherModel, error) {
	publishers, err := datatransfer.PublisherData{}.GetAll()
	if err != nil {
		return &[]entities.Publisher{}, err
	}

	var newPublisher []models.PublisherModel
	var dataPublisher []models.PublisherModel

	for _,publisher := range publishers {
		newPublisher.ID = publisher.ID
		newPublisher.PublisherName = publisher.PublisherName
		dataAuthor = append(dataPublisher, newPublisher)
	}
	return &dataPublisher, nil
}

func (p Publisher) GetById(id int) ([]models.PublisherModel, error) {
	publisher, err := datatransfer.PublisherData{}.GetById(id)
	if err != nil {
		return &[]entities.Publisher{}, err
	}

	var newPublisher []models.PublisherModel

	newPublisher.ID = publisher.ID
	newPublisher.PublisherName = publisher.PublisherName

	return &newPublisher, nil

}

func (p Publisher) Create(newPublisher *models.PublisherModel) error {
	err := datatransfer.PublisherData{}.Create(&newPublisher)
	if err != nil {
		return err
	}
	return nil
}

func (p Publisher) Edit(editedPublisher *models.PublisherModel) error {
	err := datatransfer.PublisherData{}.Edit(&editedPublisher)
	if err != nil {
		return err
	}
	return nil
}

func (p Publisher) Delete(deletedPublisher *models.PublisherModel) error {
	err := datatransfer.PublisherData{}.Delete(&deletedPublisher)
	if err != nil {
		return err
	}
	return nil
}

