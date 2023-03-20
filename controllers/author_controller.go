package controllers

import (
	// "strconv"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/aabdullahgungor/mybookcase/business"
)

type AuthorController struct{
}

func (a AuthorController) GetAll(c *gin.Context)  {
	authors, err := business.Author{}.GetAll()
	if err != nil {
		return err
	}
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