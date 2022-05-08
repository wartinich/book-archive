package handler

import "github.com/gin-gonic/gin"

func Handler() *gin.Engine {
	router := gin.Default()

	book := router.Group("/book")
	{
		book.POST("", CreateBook)
		book.GET("/list", BookList)
		book.GET("/:id", BookDetail)
	}

	author := router.Group("/author")
	{
		author.POST("", CreateAuthor)
		author.GET("/list", AuthorList)
		author.GET("/:id", AuthorDetail)
	}

	return router
}
