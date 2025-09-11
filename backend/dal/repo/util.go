package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type sqlIn interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int | ~uint | ~string | ~float32 | ~float64
}

func onlyIn[T sqlIn](db *sqlx.DB, query string, args []T) (sql.Result, error) {
	if len(args) == 0 {
		return db.Exec(query, "")
	}
	query, arg, err := sqlx.In(query, args)

	if err != nil {
		return nil, err
	}

	return db.Exec(query, arg...)
}
