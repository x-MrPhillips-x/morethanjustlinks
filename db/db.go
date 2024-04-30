package db

import (
	"example.com/morethanjustlinks/models"
	"gorm.io/gorm"
)

var SQLDSN = "root:secret@tcp(morethanjustlinks-maria-db)/morethanjustlinks_db?charset=utf8mb4&parseTime=True&loc=Local"

type SqlDB interface {
	Ping() error
}

func Ping(db SqlDB) error {
	return db.Ping()
}

func NewGormDB(dialector gorm.Dialector, cfg *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, cfg)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, err
}
