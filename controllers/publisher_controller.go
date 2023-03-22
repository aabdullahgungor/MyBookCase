package controllers

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/entities"
)

type PublisherController struct{
}

func (p PublisherController) GetAll(c *gin.Context)  {
	var publisherModel models.PublisherModel
	publishers, _ := publisherModel.GetAll()
	c.IndentedJSON(http.StatusOK, publishers)
}

func (p PublisherController) GetById(c *gin.Context)  {
	str_id := c.Param("id")
	int_id, _ := strconv.Atoi(str_id)
	var publisherModel models.PublisherModel
	publisher, _ := publisherModel.GetById(int_id)
	c.IndentedJSON(http.StatusOK, publisher)
}

func (p PublisherController) Create(c *gin.Context)  {
	
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

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Publisher has been created","publisher_id": publisher.ID})
}

func (p PublisherController) Edit(c *gin.Context)  {
	
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

	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Publisher has been edited","publisher_id": publisher.ID})
}

func (p PublisherController) Delete(c *gin.Context)  {
	
}