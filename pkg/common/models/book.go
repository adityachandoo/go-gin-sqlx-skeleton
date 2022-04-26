package models

type Book struct {
	Id          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Author      string `db:"author" json:"author"`
	Description string `db:"description" json:"description"`
}
