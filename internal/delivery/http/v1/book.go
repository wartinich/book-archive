package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wartinich/book-archive/internal/domain"
	"github.com/google/uuid"
)

func CreateBook(c *gin.Context) {
	var input domain.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New()

	book := domain.Book{Id: id.String(), Title: input.Title, Description: input.Description}
	domain.DB.Create(&book)

	c.IndentedJSON(http.StatusCreated, book)
}

func BookList(c *gin.Context) {
	var books []domain.Book

	domain.DB.Find(&books)

	c.IndentedJSON(http.StatusOK, books)
}

func BookDetail(c *gin.Context) {
	id := c.Param("id")
	var book domain.Book
  
	if err := domain.DB.Where("id = ?", id).First(&book).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
	  return
	}
  
	c.IndentedJSON(http.StatusOK, book)
}
