package models

import (
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type AuthorModel struct {
}


func (authorModel AuthorModel) GetAll() ([]entities.Author, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var authors []entities.Author
		db.Find(&authors)//db.Preload("Books").Find(&authors)
		//db.Table("author").Select("id,name").Scan(&authors)
		return authors, nil
	}
}

func (authorModel AuthorModel) GetById(id int) (entities.Author, error) {
	db, err := database.GetDB()
	if err != nil {
		return entities.Author{}, err
	} else {
		var author entities.Author
		db.Where("id = ?", id).First(&author)
		return author, nil
	}
}