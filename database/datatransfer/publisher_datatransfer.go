package datatransfer

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type PublisherData struct {
}

func (publisherData PublisherData) GetAll() ([]entities.Publisher, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var publishers []entities.Publisher
		db.Find(&publishers)
		return publishers, nil
	}
}

func (publisherData PublisherData) GetById(id int) ([]entities.Publisher, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var publisher []entities.Publisher
		db.Where("Id = ?", id).Find(&publisher)
		return publisher, nil
	}
}

func (publisherData PublisherData) Create(publisher *entities.Publisher) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&publisher)
		return nil
	}
}

func (publisherData PublisherData) Edit(publisher *entities.Publisher) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&publisher)
		return nil
	}
}

func (publisherData PublisherData) Delete(publisher *entities.Publisher) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(&publisher)
		return nil
	}
}

