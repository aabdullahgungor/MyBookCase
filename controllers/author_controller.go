package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/models"
)

type AuthorController struct{
}

func (a AuthorController) GetAll(c *gin.Context)  {
	var authorModel models.AuthorModel
	authors, _ := authorModel.GetAll()
	c.IndentedJSON(http.StatusOK, authors)
}

func (a AuthorController) GetById(c *gin.Context)  {
	
}

func (a AuthorController) Create(c *gin.Context)  {
	
}

func (a AuthorController) Edit(c *gin.Context)  {
	
}

func (a AuthorController) Delete(c *gin.Context)  {
	
}