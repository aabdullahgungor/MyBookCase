package controllers

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/models"
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
	
}

func (p PublisherController) Edit(c *gin.Context)  {
	
}

func (p PublisherController) Delete(c *gin.Context)  {
	
}