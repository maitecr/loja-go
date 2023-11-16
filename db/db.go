package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarBD() *sql.DB {
	conexao := "user=postgres dbname=loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err)
	}

	return db
}
