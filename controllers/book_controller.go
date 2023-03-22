package controllers 

import (
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
	
}

func (b BookController) Create(c *gin.Context)  {
	
}

func (b BookController) Edit(c *gin.Context)  {
	
}

func (b BookController) Delete(c *gin.Context)  {
	
}
