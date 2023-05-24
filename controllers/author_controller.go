package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)

type AuthorController struct {
}

func (a AuthorController) GetAll(c *gin.Context) {
	var authorModel models.AuthorModel
	authors, _ := authorModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, authors)
}

func (a AuthorController) GetById(c *gin.Context) {
	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)
	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	var authorModel models.AuthorModel
	author, err := authorModel.GetById(int_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, author)

}

func (a AuthorController) Create(c *gin.Context) {
	var author entities.Author
	err := c.ShouldBindJSON(&author)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var authorModel models.AuthorModel
	err = authorModel.Create(&author)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create author: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Author has been created", "author_id": author.ID})
}

func (a AuthorController) Edit(c *gin.Context) {
	var author entities.Author
	err := c.ShouldBindJSON(&author)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var authorModel models.AuthorModel
	err = authorModel.Edit(&author)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit author: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Author has been edited", "author_id": author.ID})

}

func (a AuthorController) Delete(c *gin.Context) {
	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)

	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	var authorModel models.AuthorModel
	err := authorModel.Delete(int_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete author: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Author has been deleted", "author_id": int_id})

}
