package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("pgx", "postgres://postgres:Manado@localhost:5432/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
