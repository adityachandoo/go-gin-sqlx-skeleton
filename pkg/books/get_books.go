package books

import (
	"net/http"

	"github.com/adityachandoo/go-gin-sqlx-skeleton/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetBooks(c *gin.Context) {

	var books []models.Book

	err := h.DB.Select(&books, "SELECT * FROM book")
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, &books)
}
