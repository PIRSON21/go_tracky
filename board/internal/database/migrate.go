package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
	"go.uber.org/zap"
)

type Migrate interface {
	Migration(string, *zap.Logger)
}

type MigrateService struct {
	Logger *zap.Logger
}

func (m *MigrateService) Migration(connStr string) {
	pg, err := sql.Open("pgx", connStr)
	if err != nil {
		m.Logger.Error("Не удалось подключиться к базе данных: %v\n", zap.Error(err))
		panic(err)
	}

	err = goose.Up(pg, "board/internal/migrations")
	if err != nil {
		m.Logger.Error("Не удалось применить миграцию: %v\n", zap.Error(err))
		panic(err)
	}
	return
}
