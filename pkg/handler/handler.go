package handler

import "github.com/gin-gonic/gin"

func Handler() *gin.Engine {
	router := gin.Default()

	router.POST("/book", CreateBook)
	router.GET("/books", BookList)
	router.GET("/book/:id", BookDetail)

	router.POST("/author", CreateAuthor)
	router.GET("/authors", AuthorList)
	router.GET("/author/:id", AuthorDetail)

	return router
}
