package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)

type ReaderController struct {
}

func (r ReaderController) GetAll(c *gin.Context) {
	var readerModel models.ReaderModel
	readers, _ := readerModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, readers)
}

func (r ReaderController) GetById(c *gin.Context) {
	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)
	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	var readerModel models.ReaderModel
	reader, err := readerModel.GetById(int_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, reader)
}

func (r ReaderController) Create(c *gin.Context) {

	var reader entities.Reader
	err := c.ShouldBindJSON(&reader)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		c.Abort()
		return
	}

	err = reader.HashPassword(reader.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var readerModel models.ReaderModel
	err = readerModel.Create(&reader)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create reader: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Reader has been created", "reader_id": reader.ID, "email": reader.Email, "name": reader.Name})

}

func (r ReaderController) Edit(c *gin.Context) {

	var reader entities.Reader
	err := c.ShouldBindJSON(&reader)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var readerModel models.ReaderModel
	err = readerModel.Edit(&reader)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit reader: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Reader has been edited", "reader_id": reader.ID})
}

func (r ReaderController) Delete(c *gin.Context) {

	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)

	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var readerModel models.ReaderModel

	err := readerModel.Delete(int_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete reader: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Reader has been deleted", "reader_id": int_id})
}
