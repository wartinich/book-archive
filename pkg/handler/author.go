package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wartinich/web/pkg/model"
	"net/http"
)

func CreateAuthor(c *gin.Context) {
	var newAuthor model.Author

	if err := c.BindJSON(&newAuthor); err != nil {
		return
	}

	model.Authors = append(model.Authors, newAuthor)
	c.IndentedJSON(http.StatusCreated, newAuthor)
}

func AuthorList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Authors)
}

func AuthorDetail(c *gin.Context) {
	id := c.Param("id")

	for _, author := range model.Authors {
		if author.Id == id {
			c.IndentedJSON(http.StatusOK, author)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Author": "Not exist"})
}
