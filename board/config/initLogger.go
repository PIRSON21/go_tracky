package config

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() *zap.Logger {
	var err error
	Logger = LogConfig()

	if err != nil {
		panic("Ошибка инициализации логгера: " + err.Error())
	}

	return Logger
}
