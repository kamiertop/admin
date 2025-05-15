package model

import "context"

type User struct {
	ID       uint32
	Name     string
	Password string
	Email    string
	Gender   string
}

type UserRepo interface {
	Delete(ctx context.Context, ids []int) error
}
