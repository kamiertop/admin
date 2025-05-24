package service

import (
	"context"

	"backend/dal/db"
	"backend/dal/model"
	"backend/dal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
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

func (u User) Register(ctx context.Context, user model.User) (uint32, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(res)

	return u.Repo.Register(ctx, u.db, user)
}

func (u User) Login(ctx context.Context, username, password string) error {
	pwd, err := u.Repo.Login(ctx, u.db, username)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password))
}
