package router

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"backend/api/handler"
	"backend/api/service"
)

func registerUser(group fiber.Router, logger *zap.Logger) {
	var (
		u = handler.User{
			Service: service.NewUser(logger),
		}
		ug = group.Group("user")
	)

	ug.Post("/register", u.Register)
	ug.Delete("/", u.Delete)
	ug.Post("/login", u.Login)
	ug.Get("/", u.List)
}
