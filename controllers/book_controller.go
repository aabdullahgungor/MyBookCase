package controllers 

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/models"
)

type BookController struct{
}

func (b BookController) GetAll(c *gin.Context)  {
	var bookModel models.BookModel
	books, _ := bookModel.GetAll()
	c.IndentedJSON(http.StatusOK, books)	
}

func (b BookController) GetById(c *gin.Context)  {
	str_id := c.Param("id")
	int_id, _ := strconv.Atoi(str_id)
	var bookModel models.BookModel
	book, _ := bookModel.GetById(int_id)
	c.IndentedJSON(http.StatusOK, book)
}

func (b BookController) Create(c *gin.Context)  {
	
}

func (b BookController) Edit(c *gin.Context)  {
	
}

func (b BookController) Delete(c *gin.Context)  {
	
}
