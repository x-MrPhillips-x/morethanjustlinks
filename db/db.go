package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DbInterface interface {
	Ping() error
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type RowsInterface interface {
	Close() error
	Err() error
	Next() bool
	Scan(dest ...any) error
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

func QueryRow(db DbInterface, query string, args ...any) *sql.Row {
	return db.QueryRow(query, args...)
}

// Rows interface

func Close(rows RowsInterface) error {
	return rows.Close()
}

func Err(rows RowsInterface) error {
	return rows.Err()
}

func Next(rows RowsInterface) bool {
	return rows.Next()
}

func Scan(rows RowsInterface, dest ...any) error {
	return rows.Scan()
}
