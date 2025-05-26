package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"backend/api/service"
	"backend/dal/model"
)

type User struct {
	Service service.User
}

// Delete 批量删除用户.
func (u User) Delete(ctx fiber.Ctx) error {
	var (
		req struct {
			IDs []int `json:"ids" validate:"required"`
		}
		err error
	)

	if err = ctx.Bind().JSON(&req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err = u.Service.Delete(ctx.Context(), req.IDs); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

// Register 用户注册.
func (u User) Register(ctx fiber.Ctx) error {
	var user model.User
	if err := ctx.Bind().JSON(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := u.Service.Register(ctx.Context(), user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

// Login 用户登录接口.
func (u User) Login(ctx fiber.Ctx) error {
	var (
		req struct {
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		}
		err error
	)

	if err = ctx.Bind().JSON(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = u.Service.Login(ctx.Context(), req.Username, req.Password); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}

func (u User) List(ctx fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if limit == 0 {
		limit = 10
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	count, list, err := u.Service.List(ctx.Context(), limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
		"list":  list,
	})
}
