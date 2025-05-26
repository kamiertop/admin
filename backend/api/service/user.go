package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"backend/dal/db"
	"backend/dal/model"
	"backend/dal/repo"
)

type User struct {
	Repo repo.User
	db   *pgxpool.Pool
}

func NewUser(logger *zap.Logger) User {
	return User{
		Repo: repo.User{
			Logger: logger,
		},
		db: db.DB,
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

func (u User) List(ctx context.Context, limit, offset int) (int, []model.User, error) {
	return u.Repo.List(ctx, u.db, limit, offset)
}
