package controllers

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/entities"
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

	var reader entities.Reader
	err := c.ShouldBindJSON(&reader)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var readerModel models.ReaderModel
	err = readerModel.Create(&reader)
	if err != nil {
		c.IndentedJSON(http.StatusCreated, gin.H{
			"error": "cannot create reader: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Reader has been created","reader_id": reader.ID})
	
}

func (r ReaderController) Edit(c *gin.Context)  {
	
}

func (r ReaderController) Delete(c *gin.Context)  {
	
}