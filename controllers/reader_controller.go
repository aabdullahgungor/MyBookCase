package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/models"
)

type ReaderController struct{
}

func (r ReaderController) GetAll(c *gin.Context)  {
	var readerModel models.ReaderModel
	readers, _ := readerModel.GetAll()
	c.IndentedJSON(http.StatusOK, readers)
}

func (r ReaderController) GetById(c *gin.Context)  {
	
}

func (r ReaderController) Create(c *gin.Context)  {
	
}

func (r ReaderController) Edit(c *gin.Context)  {
	
}

func (r ReaderController) Delete(c *gin.Context)  {
	
}