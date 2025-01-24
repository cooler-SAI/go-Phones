package database

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"path/filepath"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS phones (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    brand TEXT NOT NULL,
    price REAL NOT NULL
);`

func ConnectDB() (*sql.DB, error) {
	dbPath := "database/phones.db"
	absolutePath, _ := filepath.Abs(dbPath)
	log.Info().Msgf("Absolute path to database: %s", absolutePath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Info().Msg("Database connection established")
	return db, nil
}

func SetupDB(db *sql.DB) error {
	log.Info().Msg("Setting up database...")
	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Error().Err(err).Msg("Failed to execute createTableQuery")
	} else {
		log.Info().Msg("Table 'phones' created or already exists")
	}
	return err
}
