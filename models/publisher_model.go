package models

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type PublisherModel struct {
	
}

func (publisherModel PublisherModel) GetAll() ([]entities.Publisher, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var publishers []entities.Publisher
		db.Preload("Books").Find(&publishers)
		return publishers, nil
	}
}