package models

import (
	"errors"
	"reflect"

	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

var (
	ErrReaderIDIsNotValid   = errors.New("id is not valid")
	ErrReaderNameIsNotEmpty = errors.New("reader name cannot be empty")
	ErrReaderNotFound       = errors.New("the reader cannot be found")
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
	if id <= 0 || reflect.TypeOf(id).Kind() != reflect.Int {
		return entities.Reader{}, ErrReaderIDIsNotValid
	}
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

	if reader.Name == "" {
		return ErrReaderNameIsNotEmpty
	} else {
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Create(&reader)
			return nil
		}
	}

}

func (readerModel ReaderModel) Edit(reader *entities.Reader) error {

	if reader.Name == "" {
		return ErrReaderNameIsNotEmpty
	} else {
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Save(&reader)
			return nil
		}
	}

}

func (readerModel ReaderModel) Delete(id int) error {
	reader, err := readerModel.GetById(id)
	if err != nil {
		return err
	}
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	db.Delete(reader)
	return nil
}
