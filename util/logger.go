package util

import (
	"fmt"

	"go.uber.org/zap"
)

// InitLogger ...
func InitLogger() *zap.Logger {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		DisableStacktrace: true,
	}
	logger, err := cfg.Build()
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	zap.ReplaceGlobals(logger)

	return logger
}
