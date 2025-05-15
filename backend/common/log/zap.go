package log

import (
	"io"
	"os"
	"strings"

	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger = new(zap.Logger)

func Init() {
	zap.New(zapcore.NewCore(
		encoder(), writers(), level()),
		zap.AddCaller(),
		zap.AddCallerSkip(0),
	)
}

func writers() zapcore.WriteSyncer {
	ws := []zapcore.WriteSyncer{
		zapcore.AddSync(fileWriter()),
	}

	ws = append(ws, zapcore.AddSync(os.Stdout))

	return zapcore.NewMultiWriteSyncer(ws...)
}

func fileWriter() io.Writer {
	return &lumberjack.Logger{}
}

func level() zap.AtomicLevel {
	l := zap.NewAtomicLevel() // default level is InfoLevel
	switch strings.ToLower("config.Conf.Log.Level") {
	case "debug":
		l.SetLevel(zapcore.DebugLevel)
	case "info":
		l.SetLevel(zapcore.InfoLevel)
	case "warn":
		l.SetLevel(zapcore.WarnLevel)
	case "error":
		l.SetLevel(zapcore.ErrorLevel)
	case "panic":
		l.SetLevel(zapcore.PanicLevel)
	default:
		l.SetLevel(zapcore.InfoLevel)
	}

	return l
}

func encoder() zapcore.Encoder {
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

	//if config.Conf.System.Env == consts.DevelopmentMode {
	//	conf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	//	conf.ConsoleSeparator = " "
	//	conf.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	//	return zapcore.NewConsoleEncoder(conf)
	//}

	return zapcore.NewJSONEncoder(conf)
}
