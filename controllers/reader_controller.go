package controllers

import (
	"strconv"
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
	str_id := c.Param("id")
	int_id, _ := strconv.Atoi(str_id)
	var readerModel models.ReaderModel
	reader, _ := readerModel.GetById(int_id)
	c.IndentedJSON(http.StatusOK, reader)
}

func (r ReaderController) Create(c *gin.Context)  {
	
}

func (r ReaderController) Edit(c *gin.Context)  {
	
}

func (r ReaderController) Delete(c *gin.Context)  {
	
}