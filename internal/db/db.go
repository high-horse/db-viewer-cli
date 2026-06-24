package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
	_ "embed"
)

//go:embed migration.sql
var migrationSQL string 

var Conn *sql.DB

func InitDb() (*sql.DB, error) {
	var err error
	Conn, err = sql.Open("sqlite", "static/app.db")
	if err != nil {
		return  nil, err
	}

	if err := runMigration(Conn); err != nil {
		return  nil, err
	}

	return Conn, nil
}

func runMigration(db *sql.DB) error {
	_, err := db.Exec(migrationSQL)
	if err != nil {
		return  err
	}
	return  nil
}