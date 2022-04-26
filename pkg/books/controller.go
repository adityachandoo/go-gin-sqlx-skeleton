package books

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type handler struct {
	DB *sqlx.DB
}

func RegisterRoutes(r *gin.Engine, db *sqlx.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
