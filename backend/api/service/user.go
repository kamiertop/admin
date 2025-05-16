package service

import (
	"context"

	"backend/dal/db"
	"backend/dal/model"
	"backend/dal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Repo repo.User
	db   *pgxpool.Pool
}

func NewUser() User {
	return User{
		Repo: repo.User{},
		db:   db.DB,
	}
}

func (u User) Delete(ctx context.Context, ids []int) error {
	return u.Repo.Delete(ctx, u.db, ids)
}

func (u User) Create(ctx context.Context, user model.User) (uint32, error) {
	return u.Repo.Create(ctx, u.db, user)
}
