package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wartinich/book-archive/internal/domain"
)

// CreateAuthor  godoc
// @Summary      Create author
// @Description  create author
// @ID           create-author
// @Param        input body domain.Author true "list info"
// @Tags         authors
// @Accept       json
// @Produce      json
// @Router       /author [post]
func CreateAuthor(c *gin.Context) {
	var newAuthor domain.Author

	if err := c.BindJSON(&newAuthor); err != nil {
		return
	}

	domain.Authors = append(domain.Authors, newAuthor)
	c.IndentedJSON(http.StatusCreated, newAuthor)
}

// AuthorList 	 godoc
// @Summary      Author list
// @Description  author list
// @ID 			 author-list
// @Tags         authors
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Author
// @Router       /authors [get]
func AuthorList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, domain.Authors)
}

// AuthorDetail	 godoc
// @Summary      Author detail
// @ID 			 author-detail
// @Tags         authors
// @Param 		 id path string true "id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Author
// @Router       /author/:id [get]
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
