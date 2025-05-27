package middleware

import (
	"errors"

	"backend/common/errcode"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {
	var e *errcode.AppError

	if errors.As(err, &e) {
		return ctx.Status(e.Code).JSON(fiber.Map{
			"msg": e.Msg,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"msg": "后端同学写了bug~",
	})
}
