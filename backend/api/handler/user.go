package handler

import (
	"backend/api/service"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	Service service.User
}

// Delete / body: {"ids": [1,2,3]}
func (u User) Delete(ctx fiber.Ctx) error {
	var (
		req struct {
			Ids []int `json:"ids"`
		}
		err error
	)
	if err = ctx.Bind().JSON(&req); err != nil {
		return err
	}
	if err = u.Service.Delete(ctx.Context(), req.Ids); err != nil {
		return err
	}
	ctx.Status(fiber.StatusOK)
	return nil
}

func (User) Create(ctx fiber.Ctx) error {
	return nil
}
