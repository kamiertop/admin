package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func InitSqlite(fileName string, maxIdleConn, maxOpenConn int) (*sqlx.DB, error) {
	if fileName == "" {
		fileName = "./sqlite.db"
	}
	var err error
	DB, err = sqlx.Connect("sqlite3", fileName)
	if err != nil {
		return nil, err
	}
	DB.SetMaxIdleConns(maxIdleConn)
	DB.SetMaxOpenConns(maxOpenConn)

	return DB, nil
}
