package router

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

func Serve(addr string) error {
	app := fiber.New(fiber.Config{
		AppName:      "admin",
		ServerHeader: "Fiber",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	// 注册路由
	registerRoute(app)

	return app.Listen(addr, fiber.ListenConfig{
		// 使用context.WithCancel + signal实现优雅关机, 而不是使用app.Shutdown()
		GracefulContext:       shutdownCtx(),
		ShutdownTimeout:       5 * time.Second,
		DisableStartupMessage: true,
		EnablePrefork:         false,
		EnablePrintRoutes:     false,
	})
}
func shutdownCtx() context.Context {
	ctx, cancelFunc := context.WithCancel(context.Background())
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-exitCh
		cancelFunc()
	}()

	return ctx
}

func registerRoute(app *fiber.App) {
	app.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.SendString("pong")
	})
	rootGroup := app.Group("")

	registerUser(rootGroup)
}
