package handler

import (
	"backend/api/resp"

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
			IDs []int `json:"ids" validate:"required"`
		}
		err error
	)

	if err = ctx.Bind().JSON(&req); err != nil {
		return err
	}

	if err = u.Service.Delete(ctx.Context(), req.IDs); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(resp.Success(nil))
}

func (u User) Register(ctx fiber.Ctx) error {
	var user model.User
	if err := ctx.Bind().JSON(&user); err != nil {
		return err
	}

	id, err := u.Service.Register(ctx.Context(), user)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(resp.Success(fiber.Map{
		"id": id,
	}))
}

func (u User) Login(ctx fiber.Ctx) error {
	var (
		req struct {
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		}
		err error
	)
	if err = ctx.Bind().JSON(&req); err != nil {
		return err
	}
	if err = u.Service.Login(ctx.Context(), req.Username, req.Password); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(resp.Success(nil))
}
