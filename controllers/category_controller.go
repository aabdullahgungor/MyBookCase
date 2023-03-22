package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/models"
)


type CategoryController struct{
}

func (C CategoryController) GetAll(c *gin.Context)  {
	var categoryModel models.CategoryModel
	categories, _ := categoryModel.GetAll()
	c.IndentedJSON(http.StatusOK, categories)	
}

func (C CategoryController) GetById(c *gin.Context)  {
	
}

func (C CategoryController) Create(c *gin.Context)  {
	
}

func (C CategoryController) Edit(c *gin.Context)  {
	
}

func (C CategoryController) Delete(c *gin.Context)  {
	
}