package tools

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"go-Phones/handlers"
)

func StartServer() {
	log.Info().Msg("Welcome to the Phones Web Server Application!")
	log.Info().Msg("Server is running!")
	log.Info().Msg("Main page: http://localhost:8080")
	log.Info().Msg("Phone list: http://localhost:8080/phones")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.WelcomePage)         // Главная страница
	http.HandleFunc("/phones", handlers.PhoneListPage) // Страница со списком телефонов

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}
