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

// GetAuthors            godoc
// @Summary		Get authors array
// @Description	Responds with the list of all authors as JSON.
// @Tags			authors
// @Produce		json
// @Success		200	{object}	entities.Author
// @Router			/authors [get]
func (a AuthorController) GetAll(c *gin.Context) {
	var authorModel models.AuthorModel
	authors, _ := authorModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, authors)
}

// GetAuthor           godoc
// @Summary		Get single author by id
// @Description	Returns the author whose id value matches the id.
// @Tags			authors
// @Produce		json
// @Param			id path	string true "search author by id"
// @Success		200		{object}	entities.Author
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/authors/{id} [get]
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

// CreateAuthor           godoc
// @Summary		Add a new author
// @Description	Takes a author JSON and store in DB. Return saved JSON.
// @Tags			authors
// @Produce		json
// @Param			author body	entities.Author	true "Author JSON"
// @Success		200		{object}	response.Response{}
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/authors [post]
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

// EditAuthor           godoc
// @Summary		Edit an author
// @Description	Takes a author JSON and edit an author in DB. Return saved JSON.
// @Tags			authors
// @Produce		json
// @Param			author body	entities.Author	true "Author JSON"
// @Success		200		{object}	response.Response{}
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/authors [put]
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

// DeleteAuthor           godoc
// @Summary		Delete an author
// @Description	Remove an author from DB by id.
// @Tags			authors
// @Produce		json
// @Param			id path	string true "delete author by id"
// @Success		200		{object}	response.Response{}
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/authors/{id} [delete]
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
