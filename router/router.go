package router

import (
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine{
	main := router.Group("api/v1") 
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.GetAllBooks)
			books.GET("/:id", controllers.GetBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		authors := main.Group("authors")
		{
			authors.GET("/", controllers.GetAllAuthors)
			authors.GET("/:id", controllers.GetAuthor)
			authors.POST("/", controllers.CreateAuthor)
			authors.PUT("/", controllers.EditAuthor)
			authors.DELETE("/:id", controllers.DeleteAuthor)
		}
		categories := main.Group("categories")
		{
			categories.GET("/", controllers.GetAllCategories)
			categories.GET("/:id", controllers.GetCategory)
			categories.POST("/", controllers.CreateCategory)
			categories.PUT("/", controllers.EditCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
		}
		publishers := main.Group("publishers")
		{
			publishers.GET("/", controllers.GetAllPublishers)
			publishers.GET("/:id", controllers.GetPublisher)
			publishers.POST("/", controllers.CreatePublisher)
			publishers.PUT("/", controllers.EditPublisher)
			publishers.DELETE("/:id", controllers.DeletePublisher)
		}
		readers := main.Group("readers")
		{
			readers.GET("/", controllers.GetAllReaders)
			readers.GET("/:id", controllers.GetReader)
			readers.POST("/", controllers.CreateReader)
			readers.PUT("/", controllers.EditReader)
			readers.DELETE("/:id", controllers.DeleteReader)
		}

	}
}