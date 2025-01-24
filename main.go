package main

import (
	"database/sql"
	"go-Phones/commands"
	"go-Phones/database"
	"go-Phones/handlers"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

func main() {
	// Настройка Zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Проверяем аргументы командной строки
	if len(os.Args) > 1 {
		// Если переданы аргументы, запускаем CLI
		if err := commands.Execute(); err != nil {
			log.Fatal().Err(err).Msg("CLI execution error")
		}
		return
	}

	// Если аргументов нет, запускаем веб-сервер
	log.Info().Msg("Welcome to the Phones Web Server Application!")
	log.Info().Msg("Server will run at http://localhost:8080")

	// Инициализация базы данных
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

	// Настройка базы данных
	if err := database.SetupDB(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to set up database")
	}

	// Установим базу данных для обработчиков
	handlers.SetDB(db)

	// Запуск веб-сервера
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.WelcomePage)

	log.Info().Msg("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}
