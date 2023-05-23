package router

import (
	"github.com/aabdullahgungor/mybookcase/controllers"
	"github.com/aabdullahgungor/mybookcase/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.BookController{}.GetAll)
			books.GET("/:id", controllers.BookController{}.GetById)
			books.POST("/", controllers.BookController{}.Create)
			books.PUT("/", controllers.BookController{}.Edit)
			books.DELETE("/:id", controllers.BookController{}.Delete)
		}
		authors := main.Group("authors")
		{
			authors.GET("/", controllers.AuthorController{}.GetAll)
			authors.GET("/:id", controllers.AuthorController{}.GetById)
			authors.POST("/", controllers.AuthorController{}.Create)
			authors.PUT("/", controllers.AuthorController{}.Edit)
			authors.DELETE("/:id", controllers.AuthorController{}.Delete)
		}

		categories := main.Group("categories")
		{
			categories.GET("/", controllers.CategoryController{}.GetAll)
			categories.GET("/:id", controllers.CategoryController{}.GetById)
			categories.POST("/", controllers.CategoryController{}.Create)
			categories.PUT("/", controllers.CategoryController{}.Edit)
			categories.DELETE("/:id", controllers.CategoryController{}.Delete)
		}
		publishers := main.Group("publishers")
		{
			publishers.GET("/", controllers.PublisherController{}.GetAll)
			publishers.GET("/:id", controllers.PublisherController{}.GetById)
			publishers.POST("/", controllers.PublisherController{}.Create)
			publishers.PUT("/", controllers.PublisherController{}.Edit)
			publishers.DELETE("/:id", controllers.PublisherController{}.Delete)
		}
		readers := main.Group("readers")
		{
			readers.GET("/", controllers.ReaderController{}.GetAll)
			readers.GET("/:id", controllers.ReaderController{}.GetById)
			readers.POST("/", controllers.ReaderController{}.Create)
			readers.PUT("/", controllers.ReaderController{}.Edit)
			readers.DELETE("/:id", controllers.ReaderController{}.Delete)
		}

		main.POST("/token", controllers.GenerateToken)
		secured := main.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}

	}
	return router
}
