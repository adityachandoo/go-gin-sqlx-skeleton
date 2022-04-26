package books

import (
	"net/http"

	"github.com/adityachandoo/go-gin-sqlx-skeleton/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	var book = models.Book{}

	err := h.DB.Get(&book, "SELECT * FROM book WHERE id = $1", id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, &book)
}
