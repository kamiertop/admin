package handler

import (
	"github.com/gofiber/fiber/v3"

	"backend/api/service"
	"backend/dal/model"
)

type User struct {
	Service service.User
}

// Delete / body: {"ids": [1,2,3]}.
func (u User) Delete(ctx fiber.Ctx) error {
	var (
		req struct {
			IDs []int `json:"ids"`
		}
		err error
	)

	if err = ctx.Bind().JSON(&req); err != nil {
		return err
	}

	if err = u.Service.Delete(ctx.Context(), req.IDs); err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK)

	return nil
}

func (u User) Create(ctx fiber.Ctx) error {
	var user model.User
	if err := ctx.Bind().JSON(&user); err != nil {
		return err
	}

	id, err := u.Service.Create(ctx.Context(), user)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}
