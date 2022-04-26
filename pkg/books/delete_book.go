package books

import (
	"net/http"

	//"github.com/adityachandoo/go-gin-sqlx-skeleton/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	//var book models.Book

	_, err := h.DB.Exec("DELETE FROM book WHERE id = $1", id)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
}
