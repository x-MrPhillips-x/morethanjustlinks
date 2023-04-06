package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DbInterface interface {
	Ping() error
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

func Connect() (*sql.DB, error) {
	return sql.Open("mysql", "root:secret@tcp(morethanjustlinks-maria-db)/morethanjustlinks_db")
}

func Ping(db DbInterface) error {
	return db.Ping()
}

func Exec(db DbInterface, query string, args ...any) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Query(db DbInterface, query string, args ...any) (*sql.Rows, error) {
	return db.Query(query, args...)
}
