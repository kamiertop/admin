package handler

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"

	"backend/api/service"
	"backend/common/errcode"
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
		return errcode.NewBadRequestError(err)
	}

	if err = u.Service.Delete(req.IDs); err != nil {
		return errcode.NewInternalError(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ok",
	})
}

// Register 用户注册.
func (u User) Register(ctx fiber.Ctx) error {
	var user model.User
	if err := ctx.Bind().JSON(&user); err != nil {
		return errcode.NewBadRequestError(err)
	}

	id, err := u.Service.Register(user)
	if err != nil {
		return errcode.NewInternalError(err)
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

	if err = u.Service.Login(req.Username, req.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errcode.NewError(fiber.StatusUnauthorized, "用户不存在", err)
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errcode.NewError(fiber.StatusUnauthorized, "密码错误", err)
		}
		return errcode.NewInternalError(err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "success",
	})
}

func (u User) List(ctx fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		return errcode.NewBadRequestError(err)
	}

	if limit == 0 {
		limit = 10
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return errcode.NewBadRequestError(err)
	}

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	count, list, err := u.Service.List(limit, offset)
	if err != nil {
		return errcode.NewInternalError(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
		"list":  list,
	})
}
