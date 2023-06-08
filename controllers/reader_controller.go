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

// GetReaders           godoc
// @Summary		Get readers array
// @Description	Responds with the list of all readers as JSON.
// @Tags			readers
// @Produce		json
// @Success		200	{object}	entities.Reader
// @Router			/readers [get]
func (r ReaderController) GetAll(c *gin.Context) {
	var readerModel models.ReaderModel
	readers, _ := readerModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, readers)
}

// GetReader           godoc
// @Summary		Get single reader by id
// @Description	Returns the reader whose id value matches the id.
// @Tags			readers
// @Produce		json
// @Param			id path	string true "search reader by id"
// @Success		200		{object}	entities.Reader
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/readers/{id} [get]
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

// CreateReader           godoc
// @Summary		Add a new reader
// @Description	Takes a reader JSON and store in DB. Return saved JSON.
// @Tags			readers
// @Produce		json
// @Param			reader body	entities.Reader	true "Reader JSON"
// @Success		200		{object}	entities.Reader
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/readers [post]
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

// EditReader           godoc
// @Summary		Edit an reader
// @Description	Takes a reader JSON and edit an in DB. Return saved JSON.
// @Tags			readers
// @Produce		json
// @Param			reader body	entities.Reader	true "Reader JSON"
// @Success		200		{object}	entities.Reader
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/readers [put]
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

// DeleteReader          godoc
// @Summary		Delete an reader
// @Description	Remove an reader from DB by id.
// @Tags			readers
// @Produce		json
// @Param			id path	string true "delete reader by id"
// @Success		200		{object}	entities.Reader
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/readers/{id} [delete]
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
