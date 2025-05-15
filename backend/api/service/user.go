package service

import (
	"context"

	"backend/dal/model"
)

type User struct {
	Repo model.UserRepo
}

func (u User) Delete(ctx context.Context, ids []int) error {
	return u.Repo.Delete(ctx, ids)
}
