package main

import (
	"database/sql"
	"go-Phones/commands"
	"go-Phones/database"
	"go-Phones/handlers"
	"go-Phones/tools"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close database")
		}
	}(db)

	if err := database.SetupDB(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to set up database")
	}

	handlers.SetDB(db)

	go tools.StartServer()

	commands.Commands(db)

	tools.StopApp()

}
