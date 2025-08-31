package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"backend/dal/model"
)

type User struct {
	Logger *zap.Logger
}

func (User) Delete(ctx context.Context, db *sqlx.DB, ids []int) error {
	if _, err := db.ExecContext(ctx, "DELETE FROM users WHERE id = ANY($1)", ids); err != nil {
		return err
	}

	return nil
}

func (User) Register(ctx context.Context, db *sqlx.DB, user model.User) (uint32, error) {
	if err := db.QueryRowContext(ctx,
		"INSERT INTO users (email, gender, name, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Email, user.Gender, user.Name, user.Password).
		Scan(&user.ID); err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (u User) Login(ctx context.Context, db *sqlx.DB, username string) (string, error) {
	var password string
	if err := db.QueryRowContext(ctx, "SELECT password FROM users WHERE name = $1", username).Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}

func (u User) List(ctx context.Context, db *sqlx.DB, limit, offset int) (int, []model.User, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, nil, err
	}

	// defer func(tx pgx.Tx, ctx context.Context) {
	// 	// pgx.ErrTxClosed: 正常情况
	// 	if err = tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
	// 		u.Logger.Error("rollback error", zap.Error(err))
	// 	}
	// }(tx, ctx)

	rows, err := tx.Query("SELECT id, name, email, gender FROM users ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return 0, nil, err
	}

	defer rows.Close()

	res := make([]model.User, 0, limit)

	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Gender); err != nil {
			return 0, nil, err
		}

		res = append(res, user)
	}

	var count int
	if err = tx.QueryRowContext(ctx, "SELECT count(id) FROM users").Scan(&count); err != nil {
		return 0, nil, err
	}

	if err := tx.Commit(); err != nil {
		return 0, nil, err
	}

	return count, res, nil
}
