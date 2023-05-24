package router

import (
	"github.com/aabdullahgungor/mybookcase/controllers"
	"github.com/aabdullahgungor/mybookcase/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
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

		// programmatically set swagger info
		docs.SwaggerInfo.Title = "Swagger Example API"
		docs.SwaggerInfo.Description = "This is a sample server AlbumStore server."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = ""
		docs.SwaggerInfo.BasePath = ""
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		// add swagger
		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}
	return router
}
