package models

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type ReaderModel struct {
}


func (readerModel ReaderModel) GetAll() ([]entities.Reader, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var readers []entities.Reader
		db.Find(&readers)
		return readers, nil
	}
}

func (readerModel ReaderModel) GetById(id int) (entities.Reader, error) {
	db, err := database.GetDB()
	if err != nil {
		return entities.Reader{}, err
	} else {
		var reader entities.Reader
		db.Where("id = ?", id).First(&reader)
		return reader, nil
	}
}

func (readerModel ReaderModel) Create(reader *entities.Reader) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&reader)
		return nil
	}
}

func (readerModel ReaderModel) Edit(reader *entities.Reader) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&reader)
		return nil
	}
}

func (readerModel ReaderModel) Delete(reader entities.Reader) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(reader)
		return nil
	}
}