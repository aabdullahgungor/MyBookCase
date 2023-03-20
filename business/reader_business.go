package business

import (
	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/database/datatransfer"
)

type Reader struct{
}

func (r Reader) GetAll() ([]models.ReaderModel, error) {
	readers, err := datatransfer.ReaderData{}.GetAll()
	if err != nil {
		return &[]entities.Reader{}, err
	}

	var newReader []models.ReaderModel
	var dataReader []models.ReaderModel

	for _,reader := range readers {
		newReader.ID = reader.ID
		newReader.Name = reader.Name
		dataReader = append(dataReader, newReader)
	}
	return &dataReader, nil
}

func (r Reader) GetById(id int) ([]models.ReaderModel, error) {
	reader, err := datatransfer.ReaderData{}.GetById(id)
	if err != nil {
		return &[]entities.Reader{}, err
	}

	var newReader []models.ReaderModel

	newReader.ID = reader.ID
	newReader.Name = reader.Name

	return &newReader, nil

}

func (r Reader) Create(newReader *models.ReaderModel) error {
	err := datatransfer.ReaderData{}.Create(&newReader)
	if err != nil {
		return err
	}
	return nil
}

func (r Reader) Edit(editedReader *models.ReaderModel) error {
	err := datatransfer.ReaderData{}.Edit(&editedReader)
	if err != nil {
		return err
	}
	return nil
}

func (r Reader) Delete(deletedReader *models.ReaderModel) error {
	err := datatransfer.ReaderData{}.Delete(&deletedReader)
	if err != nil {
		return err
	}
	return nil
}

