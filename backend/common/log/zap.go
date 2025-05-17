package log

import (
	"io"
	"os"
	"strings"
	"time"

	"backend/config"

	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger = new(zap.Logger)

func Init(cfg config.Log) *zap.Logger {
	Logger = zap.New(zapcore.NewCore(
		encoder(cfg.Mode), writers(cfg.Mode), level(cfg.Level)),
		zap.AddCaller(),
		zap.AddCallerSkip(0),
	)

	return Logger
}

func writers(mode string) zapcore.WriteSyncer {
	mode = strings.ToLower(mode)
	if mode == "development" || mode == "dev" {
		return os.Stdout
	} else {
		return zapcore.AddSync(fileWriter())
	}
}

func fileWriter() io.Writer {
	return &lumberjack.Logger{}
}

func level(level string) zapcore.Level {
	var l zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		l = zapcore.DebugLevel
	case "info":
		l = zapcore.InfoLevel
	case "warn":
		l = zapcore.WarnLevel
	case "error":
		l = zapcore.ErrorLevel
	case "panic":
		l = zapcore.PanicLevel
	case "fatal":
		l = zapcore.FatalLevel
	default:
		l = zapcore.InfoLevel
	}

	return l
}

func encoder(mode string) zapcore.Encoder {
	conf := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
		NewReflectedEncoder: func(writer io.Writer) zapcore.ReflectedEncoder {
			return json.NewEncoder(writer)
		},
	}
	mode = strings.ToLower(mode)
	if mode == "development" || mode == "dev" {
		conf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		conf.ConsoleSeparator = " "
		conf.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
		return zapcore.NewConsoleEncoder(conf)
	} else {
		return zapcore.NewJSONEncoder(conf)
	}
}
