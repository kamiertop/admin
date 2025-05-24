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
	"github.com/gofiber/fiber/v3/middleware/pprof"
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

func Serve(addr string) error {
	app := fiber.New(fiber.Config{
		AppName:         "admin",
		ServerHeader:    "Fiber",
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		StructValidator: newValidator(),
	})

	// 注册路由
	registerRoute(app)

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

func registerRoute(app *fiber.App) {
	app.Use(pprof.New(pprof.Config{
		Next: func(ctx fiber.Ctx) bool {
			return !strings.HasPrefix(ctx.Path(), "admin")
		},
	}))

	app.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	rootGroup := app.Group("")

	registerUser(rootGroup)
}
