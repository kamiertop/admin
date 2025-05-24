package resp

import "github.com/gofiber/fiber/v3"

func Success(data any) fiber.Map {
	return fiber.Map{
		"msg":  "success",
		"data": data,
	}
}
