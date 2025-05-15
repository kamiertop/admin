package repo

import (
	"context"

	"backend/dal/db"
)

type User struct{}

func (User) Delete(ctx context.Context, ids []int) error {
	if _, err := db.DB.Exec(ctx, "DELETE FROM users WHERE id = ANY($1)", ids); err != nil {
		return err
	}

	return nil
}
