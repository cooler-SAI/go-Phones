package main

import (
	"database/sql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-Phones/handlers"
	"go-Phones/tools"
	_ "modernc.org/sqlite"
	"net/http"
	"os"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS phones (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    brand TEXT NOT NULL,
    price REAL NOT NULL
);`

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

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Server Starting....")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.WelcomePage)

	log.Info().Msg("Server StopApp Running....")
	tools.StopApp()

	log.Info().Msg("Main End..")
}
