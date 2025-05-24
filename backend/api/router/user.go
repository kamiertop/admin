package router

import (
	"github.com/gofiber/fiber/v3"

	"backend/api/handler"
	"backend/api/service"
)

func registerUser(group fiber.Router) {
	var (
		u = handler.User{
			Service: service.NewUser(),
		}
		ug = group.Group("user")
	)

	ug.Post("/register", u.Register)
	ug.Delete("/", u.Delete)
	ug.Post("/login", u.Login)
}
