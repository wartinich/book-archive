package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wartinich/book-archive/internal/domain"
)

func CreateAuthor(c *gin.Context) {
	var newAuthor domain.Author

	if err := c.BindJSON(&newAuthor); err != nil {
		return
	}

	domain.Authors = append(domain.Authors, newAuthor)
	c.IndentedJSON(http.StatusCreated, newAuthor)
}

func AuthorList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, domain.Authors)
}

func AuthorDetail(c *gin.Context) {
	id := c.Param("id")

	for _, author := range domain.Authors {
		if author.Id == id {
			c.IndentedJSON(http.StatusOK, author)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Author": "Not exist"})
}
