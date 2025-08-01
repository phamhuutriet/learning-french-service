package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Info("Hello, World!")

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "John"), zap.Int("age", 30))
	// logger.Info("Hello Example")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	// Customize logger
	encoder := getEncoderConfig()
	writer := getLogWriter()

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Hello", zap.String("name", "John"), zap.Int("age", 30))
	logger.Error("Error", zap.String("name", "John"), zap.Int("age", 30))
}

// format the log
func getEncoderConfig() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 1716931200 -> 2024-05-27T00:00:00.000Z
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encoderConfig.TimeKey = "time"
	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// /Users/mac/go/src/learning-french-service/cmd/cli/main.go:19 -> main.go:19
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// write the log to file and console
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
