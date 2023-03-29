package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)

type BookController struct{
}

func (b BookController) GetAll(c *gin.Context)  {
	var bookModel models.BookModel
	books, _ := bookModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, books)	
}

func (b BookController) GetById(c *gin.Context)  {
	str_id := c.Param("id")
	int_id, _ := strconv.Atoi(str_id)
	var bookModel models.BookModel
	book, _ := bookModel.GetById(int_id)
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, book)
}

func (b BookController) Create(c *gin.Context)  {
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

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Book has been created","book_id": book.ID})
	
}

func (b BookController) Edit(c *gin.Context)  {
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

	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Book has been edited","book_id": book.ID})
	
}

func (b BookController) Delete(c *gin.Context)  {
	
	str_id := c.Param("id")
	int_id, err := strconv.Atoi(str_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var bookModel models.BookModel
	book, _ := bookModel.GetById(int_id)

	err = bookModel.Delete(book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}


	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Book has been deleted","book_id": book.ID})
}
