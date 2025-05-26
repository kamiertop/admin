package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	var (
		e    *fiber.Error
		code = fiber.StatusInternalServerError
		msg  = "后端开发同学写了bug😡"
	)

	if errors.As(err, &e) {
		code = e.Code
	}

	if code == fiber.StatusBadRequest {
		msg = "参数有误😲"
	}

	if code == fiber.StatusInternalServerError {
		msg = "服务器开小差啦😳"
	}

	return ctx.Status(code).SendString(msg)
}
