package datatransfer

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type ReaderData struct {
}

func (readerData ReaderData) GetAll() ([]entities.Reader, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var readers []entities.Reader
		db.Find(&readers)
		return readers, nil
	}
}

func (readerData ReaderData) GetById(id int) ([]entities.Reader, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var reader []entities.Reader
		db.Where("Id = ?", id).Find(&reader)
		return reader, nil
	}
}

func (readerData ReaderData) Create(reader *entities.Reader) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&reader)
		return nil
	}
}

func (readerData ReaderData) Edit(reader *entities.Reader) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&reader)
		return nil
	}
}

func (readerData ReaderData) Delete(reader *entities.Reader) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(&reader)
		return nil
	}
}

