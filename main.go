package main

import (
	"database/sql"
	"go-Phones/handlers"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS phones (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    brand TEXT NOT NULL,
    price REAL NOT NULL
);`

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Server starting...")

	db, err := connectDB()
	if err != nil {
		log.Fatal().Err(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	if err := setupDB(db); err != nil {
		log.Fatal().Err(err)
	}

	handlers.SetDB(db)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.WelcomePage)

	log.Info().Msg("Server running at http://localhost:8080")
	err2 := http.ListenAndServe(":8080", nil)
	if err2 != nil {
		return
	}
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "phones.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func setupDB(db *sql.DB) error {
	_, err := db.Exec(createTableQuery)
	return err
}
