package config

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger = LogConfig()

	if err != nil {
		panic("Ошибка инициализации логгера: " + err.Error())
	}
}
