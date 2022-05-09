package handler

import "github.com/gin-gonic/gin"

func Handler() *gin.Engine {
	router := gin.Default()

	book := router.Group("")
	{
		book.POST("/book", CreateBook)
		book.GET("/books", BookList)
		book.GET("/book/:id", BookDetail)
	}

	author := router.Group("")
	{
		author.POST("/author", CreateAuthor)
		author.GET("/authors", AuthorList)
		author.GET("/author/:id", AuthorDetail)
	}

	return router
}
