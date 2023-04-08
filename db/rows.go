package db

import "database/sql"

type RowsInterface interface {
	Close() error
	ColumnTypes() ([]*sql.ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...any) error
}

func Next(rows RowsInterface) bool {
	return rows.Next()
}

func Scan(rows RowsInterface, dest ...any) error {
	return rows.Scan(dest...)
}
