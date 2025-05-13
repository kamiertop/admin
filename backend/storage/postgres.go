package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitPostgres(url string) (err error) {
	ctx := context.Background()
	DB, err = pgxpool.New(ctx, url)
	if err != nil {
		return err
	}

	return DB.Ping(ctx)
}
