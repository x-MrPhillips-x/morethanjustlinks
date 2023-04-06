package main

import (
	"log"

	"example.com/morethanjustlinks/db"
	"example.com/morethanjustlinks/handler"
	"go.uber.org/zap"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	s, err := handler.NewHandlerService(db, logger.Sugar(), handler.PING_DB_ATTEMPTS)
	if err != nil {
		log.Fatal(err)
	}

	router := s.SetupHandlerServiceRoutes()
	router.Run(":8080")
}
