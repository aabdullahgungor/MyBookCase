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
		db.Find(&publishers)
		return publishers, nil
	}
}

func (publisherModel PublisherModel) GetById(id int) (entities.Publisher, error) {
	db, err := database.GetDB()
	if err != nil {
		return entities.Publisher{}, err
	} else {
		var publisher entities.Publisher
		db.Where("id = ?", id).First(&publisher)
		return publisher, nil
	}
}