package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)

type PublisherController struct {
}

// GetPublishers            godoc
// @Summary		Get publishers array
// @Description	Responds with the list of all publishers as JSON.
// @Tags			publishers
// @Produce		json
// @Success		200	{object}	entities.Publisher
// @Router			/publishers [get]
func (p PublisherController) GetAll(c *gin.Context) {
	var publisherModel models.PublisherModel
	publishers, _ := publisherModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, publishers)
}

// GetPublisher           godoc
// @Summary		Get single publisher by id
// @Description	Returns the publisher whose id value matches the id.
// @Tags			publishers
// @Produce		json
// @Param			id path	string true "search publisher by id"
// @Success		200		{object}	entities.Publisher
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/publishers/{id} [get]
func (p PublisherController) GetById(c *gin.Context) {
	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)
	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	var publisherModel models.PublisherModel
	publisher, err := publisherModel.GetById(int_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, publisher)
}

// CreatePublisher           godoc
// @Summary		Add a new publisher
// @Description	Takes a publisher JSON and store in DB. Return saved JSON.
// @Tags			publishers
// @Produce		json
// @Param			publisher body	entities.Publisher	true "Publisher JSON"
// @Success		200		{object}	entities.Publisher
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/publishers [post]
func (p PublisherController) Create(c *gin.Context) {

	var publisher entities.Publisher
	err := c.ShouldBindJSON(&publisher)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var publisherModel models.PublisherModel
	err = publisherModel.Create(&publisher)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit publisher: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Publisher has been created", "publisher_id": publisher.ID})
}

// EditPublisher          godoc
// @Summary		Edit an publisher
// @Description	Takes a publisher JSON and edit an in DB. Return saved JSON.
// @Tags			publishers
// @Produce		json
// @Param			publisher body	entities.Publisher	true "Publisher JSON"
// @Success		200		{object}	entities.Publisher
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/publishers [put]
func (p PublisherController) Edit(c *gin.Context) {

	var publisher entities.Publisher
	err := c.ShouldBindJSON(&publisher)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var publisherModel models.PublisherModel
	err = publisherModel.Edit(&publisher)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit publisher: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Publisher has been edited", "publisher_id": publisher.ID})
}

// DeletePublisher           godoc
// @Summary		Delete an publisher
// @Description	Remove an publisher from DB by id.
// @Tags			publishers
// @Produce		json
// @Param			id path	string true "delete publisher by id"
// @Success		200		{object}	entities.Publisher
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/publishers/{id} [delete]
func (p PublisherController) Delete(c *gin.Context) {

	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)

	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var publisherModel models.PublisherModel

	err := publisherModel.Delete(int_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete publisher: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Publisher has been deleted", "publisher_id": int_id})

}
