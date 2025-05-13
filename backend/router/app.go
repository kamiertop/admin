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

	app.Get("/ping", func(ctx fiber.Ctx) error {
		return ctx.JSON("ping", fiber.MIMETextPlainCharsetUTF8)
	})

	// 注册路由
	registerRoute(app)

	ctx, cancelFunc := context.WithCancel(context.Background())
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-exitCh
		cancelFunc()
	}()

	return app.Listen(addr, fiber.ListenConfig{
		// 使用context.WithCancel + signal实现优雅关机, 而不是使用app.Shutdown()
		GracefulContext:       ctx,
		ShutdownTimeout:       5 * time.Second,
		DisableStartupMessage: true,
		EnablePrefork:         false,
		EnablePrintRoutes:     false,
	})
}

func registerRoute(app *fiber.App) {

}
