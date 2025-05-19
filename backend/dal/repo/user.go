package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"backend/dal/model"
)

type User struct{}

func (User) Delete(ctx context.Context, db *pgxpool.Pool, ids []int) error {
	if _, err := db.Exec(ctx, "DELETE FROM users WHERE id = ANY($1)", ids); err != nil {
		return err
	}

	return nil
}

func (User) Create(ctx context.Context, db *pgxpool.Pool, user model.User) (uint32, error) {
	if err := db.QueryRow(ctx,
		"INSERT INTO users (email, gender, name, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Email, user.Gender, user.Name, user.Password).
		Scan(&user.ID); err != nil {
		return 0, err
	}

	return user.ID, nil
}
