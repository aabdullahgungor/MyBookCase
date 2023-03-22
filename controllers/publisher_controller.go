package controllers

import (
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
	
}

func (p PublisherController) Create(c *gin.Context)  {
	
}

func (p PublisherController) Edit(c *gin.Context)  {
	
}

func (p PublisherController) Delete(c *gin.Context)  {
	
}