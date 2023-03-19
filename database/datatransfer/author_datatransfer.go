package datatransfer

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type AuthorData struct {
}

func (authorData AuthorData) GetAll() ([]entities.Author, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var authors []entities.Author
		db.Find(&authors )
		return authors, nil
	}
}

func (authorData AuthorData) GetById(id int) ([]entities.Author, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var author []entities.Author
		db.Where("Id = ?", id).Find(&author)
		return author, nil
	}
}

func (authorData AuthorData) Create(author *entities.Author ) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&author)
		return nil
	}
}

func (authorData AuthorData) Edit(author *entities.Author ) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&author)
		return nil
	}
}

func (authorData AuthorData) Delete(author *entities.Author ) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(&author)
		return nil
	}
}