package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)


type CategoryController struct{
}

func (C CategoryController) GetAll(c *gin.Context)  {
	var categoryModel models.CategoryModel
	categories, _ := categoryModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, categories)	
}

func (C CategoryController) GetById(c *gin.Context)  {
	str_id := c.Param("id")
	int_id, _ := strconv.Atoi(str_id)
	var categoryModel models.CategoryModel
	category, _ := categoryModel.GetById(int_id)
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, category)
}

func (C CategoryController) Create(c *gin.Context)  {
	var category entities.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var categoryModel models.CategoryModel
	err = categoryModel.Create(&category)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create category: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Category has been created","category_id": category.ID})
	
}

func (C CategoryController) Edit(c *gin.Context)  {
	var category entities.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var categoryModel models.CategoryModel
	err = categoryModel.Edit(&category)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit category: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Category has been edited","category_id": category.ID})
	
}

func (C CategoryController) Delete(c *gin.Context)  {
	
	str_id := c.Param("id")
	int_id, err := strconv.Atoi(str_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var categoryModel models.CategoryModel
	category, _ := categoryModel.GetById(int_id)

	err = categoryModel.Delete(category)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete category: " + err.Error(),
		})
		return
	}


	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Category has been deleted","category_id": category.ID})
}