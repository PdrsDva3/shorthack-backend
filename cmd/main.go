package main

import (
	"shorthack_backend/internal/delivery"
	"shorthack_backend/pkg/config"
	"shorthack_backend/pkg/db"
	"shorthack_backend/pkg/log"
)

func main() {
	config.InitConfig()

	database := db.ConnectDB()

	logger, file1, file2 := log.InitLogger()

	defer file1.Close()
	defer file2.Close()

	delivery.Start(database, logger)
}
