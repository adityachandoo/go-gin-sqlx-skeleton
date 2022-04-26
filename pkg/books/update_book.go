package books

import (
	"net/http"
	"strconv"

	"github.com/adityachandoo/go-gin-sqlx-skeleton/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var body = models.Book{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Id, _ = strconv.Atoi(id)
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	_, err := h.DB.NamedExec(`UPDATE book SET title = :title, author = :author, description = :description WHERE id = :id`, book)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, &book)
}
