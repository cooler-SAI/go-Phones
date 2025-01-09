package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-Phones/handlers"
	"go-Phones/tools"
	"net/http"
	"os"
)

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
