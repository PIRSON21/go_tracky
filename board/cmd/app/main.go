package main

import (
	"fmt"
	"go_tracky/board/config"
	"go_tracky/board/internal/database"
	"os"
)

func main() {
	logger := config.InitLogger()
	dbpool := config.ConnectDatabase()

	migrate := database.MigrateService{logger}
	migrate.Migration(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_SSLMODE"),
	))

	defer dbpool.Close()
	defer config.Logger.Sync()
}
