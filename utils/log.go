package utils

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the application logger
var Logger *zap.Logger

// Sugar is the sugar version of the logger
var Sugar *zap.SugaredLogger

func logger() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig = zap.NewProductionEncoderConfig()
	loggerConfig.Encoding = "console"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	loggerConfig.OutputPaths = []string{"stdout"}
	loggerConfig.ErrorOutputPaths = []string{"stderr"}

	switch strings.ToLower(viper.GetString("logging.level")) {
	case "debug":
		loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "panic":
		loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	default:
		loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal("Unable to create application logger: ", err)
		panic(err)
	}

	Logger = logger
	Sugar = logger.Sugar()
}
