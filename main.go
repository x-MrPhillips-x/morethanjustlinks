package main

import (
	"log"

	"example.com/morethanjustlinks/config"
	"example.com/morethanjustlinks/db"
	"example.com/morethanjustlinks/handler"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	appConfig, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("🤬 something went wrong, creating new app config", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := db.NewGormDB(mysql.Open(appConfig.DB.SQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	s, err := handler.NewHandler(appConfig, gormDB, logger.Sugar(), handler.PING_DB_ATTEMPTS)
	if err != nil {
		log.Fatal(err)
	}

	router := s.SetupHandlerRoutes()
	router.Run(":8080")
}
