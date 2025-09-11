package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger(logger *zap.Logger) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		s := time.Now()
		err := ctx.Next()
		msg := []zapcore.Field{
			zap.String("method", ctx.Method()),
			zap.String("route", ctx.Route().Path),
			zap.String("path", ctx.Path()),
			zap.String("ip", ctx.IP()),
			zap.String("query_row", ctx.Request().URI().QueryArgs().String()),
			zap.Any("query", ctx.Queries()),
			zap.Float64("duration", time.Since(s).Seconds()),
		}

		if err != nil {
			msg = append(msg, zap.Error(err))
			logger.Error("HTTP", msg...)
		} else {
			logger.Info("HTTP", msg...)
		}

		return err
	}
}
