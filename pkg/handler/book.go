package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wartinich/web/pkg/model"
	"net/http"
)

// CreateBook godoc
// @Summary      Create book
// @Description  create book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Book
// @Router       /book [post]
func CreateBook(c *gin.Context) {
	var newBook model.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	model.Books = append(model.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func BookList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Books)
}

func BookDetail(c *gin.Context) {
	id := c.Param("id")

	for _, book := range model.Books {
		if book.Id == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Book": "Not exist"})
}
