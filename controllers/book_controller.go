package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)

type BookController struct {
}

// GetBooks            godoc
// @Summary		Get books array
// @Description	Responds with the list of all books as JSON.
// @Tags			books
// @Produce		json
// @Success		200	{object}	entities.Book
// @Router			/books [get]
func (b BookController) GetAll(c *gin.Context) {
	var bookModel models.BookModel
	books, _ := bookModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, books)
}

// GetBook          godoc
// @Summary		Get single book by id
// @Description	Returns the book whose id value matches the id.
// @Tags			books
// @Produce		json
// @Param			id path	string true "search book by id"
// @Success		200		{object}	entities.Book
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/books/{id} [get]
func (b BookController) GetById(c *gin.Context) {
	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)
	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var bookModel models.BookModel
	book, err := bookModel.GetById(int_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, book)
}

// CreateBook           godoc
// @Summary		Add a new book
// @Description	Takes a book JSON and store in DB. Return saved JSON.
// @Tags			books
// @Produce		json
// @Param			book body	entities.Book	true "Book JSON"
// @Success		200		{object}	entities.Book
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/books [post]
func (b BookController) Create(c *gin.Context) {
	var book entities.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var bookModel models.BookModel
	err = bookModel.Create(&book)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book has been created", "book_id": book.ID})

}

// EditBook          godoc
// @Summary		Edit an book
// @Description	Takes a book JSON and edit an in DB. Return saved JSON.
// @Tags			books
// @Produce		json
// @Param			book body	entities.Book	true "Book JSON"
// @Success		200		{object}	entities.Book
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/books [put]
func (b BookController) Edit(c *gin.Context) {
	var book entities.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var bookModel models.BookModel
	err = bookModel.Edit(&book)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Book has been edited", "book_id": book.ID})

}

// DeleteBook           godoc
// @Summary		Delete an book
// @Description	Remove an book from DB by id.
// @Tags			books
// @Produce		json
// @Param			id path	string true "delete book by id"
// @Success		200		{object}	entities.Book
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/books/{id} [delete]
func (b BookController) Delete(c *gin.Context) {

	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)

	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var bookModel models.BookModel

	err := bookModel.Delete(int_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Book has been deleted", "book_id": int_id})
}
