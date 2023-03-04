package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wartinich/book-archive/internal/domain"
)

func CreateBook(c *gin.Context) {
	var newBook domain.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	domain.Books = append(domain.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func BookList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, domain.Books)
}

func BookDetail(c *gin.Context) {
	id := c.Param("id")

	for _, book := range domain.Books {
		if book.Id == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Book": "Not exist"})
}
