package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
}

// GetCategories           godoc
// @Summary		Get categories array
// @Description	Responds with the list of all categories as JSON.
// @Tags			categories
// @Produce		json
// @Success		200	{object}	entities.Category
// @Router			/categories [get]
func (C CategoryController) GetAll(c *gin.Context) {
	var categoryModel models.CategoryModel
	categories, _ := categoryModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, categories)
}

// GetCategory           godoc
// @Summary		Get single category by id
// @Description	Returns the category whose id value matches the id.
// @Tags			categories
// @Produce		json
// @Param			id path	string true "search category by id"
// @Success		200		{object}	entities.Category
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/categories/{id} [get]
func (C CategoryController) GetById(c *gin.Context) {
	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)
	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var categoryModel models.CategoryModel
	category, err := categoryModel.GetById(int_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, category)
}

// CreateCategory           godoc
// @Summary		Add a new category
// @Description	Takes a category JSON and store in DB. Return saved JSON.
// @Tags			categories
// @Produce		json
// @Param			category body	entities.Category	true "Category JSON"
// @Success		200		{object}	entities.Category
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/categories [post]
func (C CategoryController) Create(c *gin.Context) {
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

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Category has been created", "category_id": category.ID})

}

// EditCategory          godoc
// @Summary		Edit an category
// @Description	Takes a category JSON and edit an in DB. Return saved JSON.
// @Tags			categories
// @Produce		json
// @Param			category body	entities.Category	true "Category JSON"
// @Success		200		{object}	entities.Category
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/categories [put]
func (C CategoryController) Edit(c *gin.Context) {
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

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Category has been edited", "category_id": category.ID})

}

// DeleteCategory          godoc
// @Summary		Delete an category
// @Description	Remove an category from DB by id.
// @Tags			categories
// @Produce		json
// @Param			id path	string true "delete category by id"
// @Success		200		{object}	entities.Category
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/categories/{id} [delete]
func (C CategoryController) Delete(c *gin.Context) {

	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)

	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var categoryModel models.CategoryModel

	err := categoryModel.Delete(int_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete category: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Category has been deleted", "category_id": int_id})
}
