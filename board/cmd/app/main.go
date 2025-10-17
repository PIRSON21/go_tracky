package main

import (
	"fmt"
	"github.com/ChanKachan/go_tracky/config"
	"github.com/ChanKachan/go_tracky/internal/database"
	"github.com/ChanKachan/go_tracky/internal/handlers"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	logger := config.InitLogger()
	dbpool := config.ConnectDatabase()

	var migrate database.Migrate = &database.MigrateService{logger}

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

	r := gin.Default()

	r.GET("/api/board/:id", handlers.GetBoardInfoById(logger, dbpool))
	r.POST("/api/create_board", handlers.PostBoard(logger, dbpool))

	r.Run(":8080")
}
