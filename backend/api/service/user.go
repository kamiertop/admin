package service

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"backend/dal/db"
	"backend/dal/model"
	"backend/dal/repo"
)

type User struct {
	Repo repo.User
	db   *sqlx.DB
}

func NewUser(logger *zap.Logger) User {
	return User{
		Repo: repo.User{
			Logger: logger,
		},
		db: db.DB,
	}
}

func (u User) Delete(ids []int) error {
	return u.Repo.Delete(u.db, ids)
}

func (u User) Register(user model.User) (uint32, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user.Password = string(res)

	return u.Repo.Register(u.db, user)
}

func (u User) Login(username, password string) error {
	pwd, err := u.Repo.Login(u.db, username)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password))
}

func (u User) List(limit, offset int) (int, []model.User, error) {
	return u.Repo.List(u.db, limit, offset)
}
