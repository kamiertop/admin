package router

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/pprof"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/zap"

	"backend/common/middleware"
)

type structValidator struct {
	validate *validator.Validate
}

func (v structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func newValidator() structValidator {
	return structValidator{
		validate: validator.New(),
	}
}

func Serve(addr string, logger *zap.Logger) error {
	app := fiber.New(fiber.Config{
		AppName:         "admin",
		ServerHeader:    "Fiber",
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		StructValidator: newValidator(),
		ErrorHandler:    middleware.ErrorHandler,
	})

	// 注册路由
	registerRoute(app, logger)

	const maxWaitTimeout = 5 * time.Second

	return app.Listen(addr, fiber.ListenConfig{
		// 使用context.WithCancel + signal实现优雅关机, 而不是使用app.Shutdown()
		GracefulContext:       shutdownCtx(),
		ShutdownTimeout:       maxWaitTimeout,
		DisableStartupMessage: true,
		EnablePrefork:         false,
		EnablePrintRoutes:     false,
	})
}

func shutdownCtx() context.Context {
	ctx, cancelFunc := context.WithCancel(context.Background())
	exitCh := make(chan os.Signal, 1)
	// ctrl+c , kill
	signal.Notify(exitCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-exitCh
		cancelFunc()
	}()

	return ctx
}

func registerRoute(app *fiber.App, logger *zap.Logger) {
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(middleware.Logger(logger))
	app.Use(pprof.New(pprof.Config{
		Next: func(ctx fiber.Ctx) bool {
			return !strings.HasPrefix(ctx.Path(), "admin")
		},
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"},
	}))

	app.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	rootGroup := app.Group("")

	registerUser(rootGroup, logger)
}
