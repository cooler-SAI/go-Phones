package main

import (
	"database/sql"
	"go-Phones/database"
	"go-Phones/handlers"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Server starting...")

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	if err := database.SetupDB(db); err != nil {
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
