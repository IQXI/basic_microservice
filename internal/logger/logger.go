package logger

import (
	conf "basic_microservice/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func GetLogger(vp conf.ConfigStruct) *zap.Logger {

	var level zapcore.Level
	switch vp.Logger.Level {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.DebugLevel
	}

	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: vp.Logger.Outputs,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger
}
