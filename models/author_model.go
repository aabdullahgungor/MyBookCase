package models

import (
	"errors"
	"reflect"

	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

var (
	ErrAuthorIDIsNotValid    = errors.New("id is not valid")
	ErrAuthorNameIsNotEmpty = errors.New("Author name cannot be empty")
	ErrAuthorNotFound   = errors.New("the author cannot be found")
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

	if id <= 0 || reflect.TypeOf(id).Kind() != reflect.Int{
		return entities.Author{}, ErrAuthorIDIsNotValid
	}

	db, err := database.GetDB()
	if err != nil {
		return entities.Author{}, err
	} 
	var author entities.Author
	db.Where("id = ?", id).First(&author)
	return author, nil
	
}


func (authorModel AuthorModel) Create(author *entities.Author) error {

	switch  {
	case author.ID <= 0 || reflect.TypeOf(author.ID).Kind() != reflect.Int :
		return ErrAuthorIDIsNotValid
	case author.Name == "":
		return ErrAuthorNameIsNotEmpty
	default:
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Create(&author)
			return nil
		}

	}	
}

func (authorModel AuthorModel) Edit(author *entities.Author) error {
	
	switch  {
	case author.ID <= 0 || reflect.TypeOf(author.ID).Kind() != reflect.Int:
		return ErrAuthorIDIsNotValid
	case author.Name == "":
		return ErrAuthorNameIsNotEmpty
	default:
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Save(&author)
			return nil
		}

	}	
	
}

func (authorModel AuthorModel) Delete(id int) error {
	
	author, err := authorModel.GetById(id)
	if err != nil {
		return err
	}
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	db.Delete(author)
	return nil
}