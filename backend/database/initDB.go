package database

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Database *sql.DB
}

func InitDB(nameDB string) (*Database, error) {
	db, err := sql.Open("sqlite3", nameDB)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (db *Database) ReadSQL(filename string) error {
	schema, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	statements := strings.Split(string(schema), ";")
	for _, st := range statements {
		_, err := db.Database.Exec(st + ";")
		if err != nil {
			return err
		}

	}

	// err = db.Triggers()
	// if err != nil {
	// 	return err
	// }

	return nil
}
