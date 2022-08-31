package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wartinich/book-archive/internal/domain"
)

// CreateBook    godoc
// @Summary      Create book
// @Description  create book
// @ID           create-book
// @Param        input body model.Book true "list info"
// @Tags         books
// @Accept       json
// @Produce      json
// @Router       /book [post]
func CreateBook(c *gin.Context) {
	var newBook domain.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	domain.Books = append(domain.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// BookList 	 godoc
// @Summary      Book list
// @Description  book list
// @ID 			 book-list
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Book
// @Router       /books [get]
func BookList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, domain.Books)
}

// BookDetail	 godoc
// @Summary      Book detail
// @Description  book detail
// @ID 			 book-detail
// @Tags         books
// @Param 		 id path string true "id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Book
// @Router       /book/:id [get]
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
