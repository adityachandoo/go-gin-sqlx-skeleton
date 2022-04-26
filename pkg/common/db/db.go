package db

import (
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func Init(url string) *sqlx.DB {

	db, err := sqlx.Connect("pgx", url)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
