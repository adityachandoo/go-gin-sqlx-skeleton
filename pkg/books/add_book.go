package books

import (
	"net/http"

	"github.com/adityachandoo/go-gin-sqlx-skeleton/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) AddBook(c *gin.Context) {
	var body = models.Book{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	_, err := h.DB.NamedExec(`INSERT INTO book (title, author, description) VALUES (:title, :author, :description)`, book)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// https://knasmueller.net/golang-using-lastinsertid-with-postgres

	c.JSON(http.StatusCreated, &book)
}
