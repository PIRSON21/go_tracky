package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Настройка выводит: уровень, время, директорию файла, сообщение
func LogConfig() *zap.Logger {
	config := zap.NewProductionEncoderConfig()

	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.TimeKey = "time"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.CallerKey = "caller"
	config.EncodeCaller = zapcore.ShortCallerEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	), zap.AddCaller())

	return logger
}
