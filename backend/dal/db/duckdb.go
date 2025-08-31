package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/marcboeker/go-duckdb/v2"
)

var DB *sqlx.DB

func InitDuckDB(dbFileName string) (*sqlx.DB, error) {
	if dbFileName == "" {
		dbFileName = ":memory:"
	}
	var err error

	DB, err = sqlx.Connect("duckdb", dbFileName)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
