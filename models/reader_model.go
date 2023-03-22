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
		db.Preload("Books").Find(&readers)
		return readers, nil
	}
}