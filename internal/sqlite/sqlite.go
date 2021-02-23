package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteHandler(dbFilename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFilename)
	if err != nil {
		return db, err
	}
	return db, nil
}

func PingDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
