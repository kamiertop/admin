package router

import (
	"backend/api/handler"
	"backend/api/service"

	"github.com/gofiber/fiber/v3"
)

func registerUser(group fiber.Router) {
	var (
		u = handler.User{
			Service: service.NewUser(),
		}
		ug = group.Group("user")
	)

	ug.Post("/", u.Create)
	ug.Delete("/", u.Delete)
}
