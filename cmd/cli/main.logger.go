package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// // Sugar
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "WOrld", 10) // like Printf
	//
	// // Logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "GOlang"), zap.Int("age", 27))

	//  // 2
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")
	//
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")
	//
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3. Custome Logger
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core)

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// "level":"info","ts":1740582833.1921015,"caller":"cli/main.logger.go:21","msg":"Hello NewProduction"}

	// ts -> 'Lúc', timestamp => datetime
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "Lúc:"

	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	os.Mkdir("./log", os.ModePerm)
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
