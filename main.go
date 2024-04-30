package main

import (
	"log"

	"example.com/morethanjustlinks/db"
	"example.com/morethanjustlinks/handler"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := db.NewGormDB(mysql.Open(db.SQLDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	s, err := handler.NewHandler(gormDB, logger.Sugar(), handler.PING_DB_ATTEMPTS)
	if err != nil {
		log.Fatal(err)
	}

	router := s.SetupHandlerRoutes()
	router.Run(":8080")
}
