package log

import (
	"go.uber.org/zap"
)

var (
	ErrorLogger *zap.Logger
	StdOutLogger *zap.Logger
)

func Error(message string) {
	ErrorLogger.Error(message)
}

func Panic(message string) {
	ErrorLogger.Panic(message)
}

func Fatal(message string) {
	ErrorLogger.Fatal(message)
}

func Info(message string) {
	StdOutLogger.Info(message)
}

func InitLoggers() error {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"spacetraders.log", "stderr"}
	ErrorLogger, err = cfg.Build()
	if err != nil {
		return err
	}

	cfg.OutputPaths = []string{"spacetraders.log", "stdout"}
	StdOutLogger, err = cfg.Build()
	return err
}