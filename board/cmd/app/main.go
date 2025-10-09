package main

import (
	"go_tracky/board/config"
)

func main() {
	config.InitLogger()
	dbpool := config.ConnectDatabase()

	defer dbpool.Close()
	defer config.Logger.Sync()
}
