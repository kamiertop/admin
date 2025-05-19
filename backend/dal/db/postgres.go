package db

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitPostgres(url string) error {
	var (
		ctx = context.Background()
		err error
	)

	DB, err = pgxpool.New(ctx, url)
	if err != nil {
		return err
	}

	return DB.Ping(ctx)
}

// customSearchPath 手动设置多个搜索路径.
func customSearchPath(url string, schemas []string) error { //nolint:unused
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return err
	}

	config.ConnConfig.RuntimeParams["search_path"] = strings.Join(schemas, ",")

	DB, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}

	return DB.Ping(ctx)
}
