package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wartinich/book-archive/internal/delivery/http/v1"
)

func Handler() *gin.Engine {
	router := gin.Default()

	book := router.Group("")
	{
		book.POST("/books", v1.CreateBook)
		book.GET("/books", v1.BookList)
		book.GET("/book/:id", v1.BookDetail)
	}

	author := router.Group("")
	{
		author.POST("/authors", v1.CreateAuthor)
		author.GET("/authors", v1.AuthorList)
		author.GET("/authors/:id", v1.AuthorDetail)
	}

	return router
}
